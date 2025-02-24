package process

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strings"

	"wsw/backend/domain/dto"
	"wsw/backend/domain/image/path"
	"wsw/backend/domain/image/url"
	"wsw/backend/ent"
	"wsw/backend/ent/types"
	"wsw/backend/lib/utils"
	"wsw/backend/model/image"
)

type (
	Input struct {
		ImageName string
		Processes []types.ImageProcess
	}
	RunnerResult struct {
		Input      Input
		PrefixPath string
		ImageURL   *string
		Status     types.StatusEnum
		Error      error
	}
	ProcessRunner interface {
		Start(*ent.Image, []types.ImageProcess) (*RunnerResult, error)
	}
	processRunnerimpl struct {
		pathProvider   path.PathProvider
		pathGenerator  path.FilenameGenerator
		urlProvider    url.Provider
		processesModel image.ImageProcesses
	}
	inputArg struct {
		Image     *ent.Image
		Processes []types.ImageProcess
	}
)

func NewProcessRunner(pathProvider path.PathProvider, pathGenerator path.FilenameGenerator, urlProvider url.Provider, processesModel image.ImageProcesses) ProcessRunner {
	return processRunnerimpl{pathProvider: pathProvider, pathGenerator: pathGenerator, urlProvider: urlProvider, processesModel: processesModel}
}

// Start implements ProcessRunner.
func (p processRunnerimpl) Start(image *ent.Image, processes []types.ImageProcess) (*RunnerResult, error) {
	inputArgs := inputArg{Image: image, Processes: processes}

	processList, err := p.createProcessList(processes)
	if err != nil {
		return nil, err
	}
	processEntity, err := p.getProcessEntity(image.ID, processList)
	if err != nil {
		return nil, err
	}
	if processEntity != nil {
		return p.createExistResult(processEntity, inputArgs)
	} else {
		processEntity, err := p.doProcess(image, processes, processList)
		if err != nil {
			return nil, err
		}
		updatedProcessEntity, err := p.processesModel.Update(processEntity, processEntity.PathPrefix, processEntity.Status, processEntity.Error)
		if err != nil {
			return nil, err
		}
		return p.createExistResult(updatedProcessEntity, inputArgs)
	}
}

func (p processRunnerimpl) doProcess(image *ent.Image, processes []types.ImageProcess, processList []Process) (*ent.ImageProcess, error) {
	hash := p.getProcessesHash(processList)
	processEntity, err := p.createProcessEntity(image, processes, hash)
	if err != nil {
		return nil, err
	}
	imagePath, err := p.runProcesses(image, processList)
	if err != nil {
		processEntity.Status = types.Error
		processEntity.Error = err.Error()

	} else {
		processEntity.PathPrefix = imagePath.RelativeDirectory
		processEntity.Status = types.Success
	}
	return processEntity, nil
}

func (p processRunnerimpl) createProcessEntity(image *ent.Image, processes []types.ImageProcess, hash string) (*ent.ImageProcess, error) {
	return p.processesModel.Create(image, processes, hash)
}

func (p processRunnerimpl) runProcesses(image *ent.Image, processList []Process) (*path.PathData, error) {
	var imagePath *path.PathData
	imagePath = p.pathProvider.Provide(image.DestinationPath, image.Filename)
	for _, process := range processList {
		prefix := process.GeneratePathPrefix()
		newPath := p.pathGenerator.GenerateFilepath(&prefix)
		toPath := p.pathProvider.Provide(newPath, image.Filename)

		processError := process.Run(*imagePath, *toPath)

		if processError != nil {
			return nil, processError
		}

		imagePath = toPath
	}
	return imagePath, nil
}

func (p processRunnerimpl) createExistResult(processEntity *ent.ImageProcess, inputArgs inputArg) (*RunnerResult, error) {
	url := p.urlProvider.Provide(dto.NewImage(inputArgs.Image.Filename, processEntity.PathPrefix))

	var errorResult error
	if processEntity.Error != "" {
		errorResult = errors.New(processEntity.Error)
	}

	return &RunnerResult{
		Input: Input{
			ImageName: inputArgs.Image.Filename,
			Processes: inputArgs.Processes,
		},
		PrefixPath: processEntity.PathPrefix,
		ImageURL:   &url,
		Status:     processEntity.Status,
		Error:      errorResult,
	}, nil
}

func (p processRunnerimpl) getProcessEntity(imageID int, processes []Process) (*ent.ImageProcess, error) {
	return p.processesModel.TryGetByHash(imageID, p.getProcessesHash(processes))
}

func (p processRunnerimpl) getProcessesHash(processList []Process) string {
	var sb strings.Builder
	for _, process := range processList {
		sb.WriteString(process.GeneratePathPrefix())
	}
	return utils.GetMD5Hash(sb.String())
}

func (p processRunnerimpl) createProcessList(processes []types.ImageProcess) ([]Process, error) {
	processList := make([]Process, 0, len(processes))
	for _, processInput := range processes {
		processFactory, err := GetProcessFactory(processInput.Type)
		if err != nil {
			return nil, err
		}
		process, err := processFactory.Create(processInput.Options)
		if err != nil {
			return nil, err
		}
		processList = append(processList, process)
	}
	return processList, nil
}

func (p processRunnerimpl) getMd5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
