package image

import (
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
	"wsw/backend/ent/types"

	"github.com/xorcare/pointer"
)

type (
	ImageProcesses interface {
		Create(*ent.UploadImage, []types.ImageProcess) (*ent.ImageProcess, error)
		Update(*ent.ImageProcess, types.StatusEnum, error) (*ent.ImageProcess, error)
	}
	imageProcessesImpl struct {
		processRepository repository.ImageProcessRepository
		imageRepository   repository.UploadImageRepository
	}
)

func (i imageProcessesImpl) Update(processEntity *ent.ImageProcess, status types.StatusEnum, err error) (*ent.ImageProcess, error) {
	var errorMessage *string
	if err != nil {
		errorMessage = pointer.String(err.Error())
	} else {
		errorMessage = nil
	}
	return i.processRepository.Update(processEntity, status, errorMessage)
}

func NewImageProcesses(processRepository repository.ImageProcessRepository, imageRepository repository.UploadImageRepository) ImageProcesses {
	return imageProcessesImpl{processRepository: processRepository, imageRepository: imageRepository}
}

func (i imageProcessesImpl) Create(imageEntity *ent.UploadImage, imageProcesses []types.ImageProcess) (*ent.ImageProcess, error) {
	return i.imageRepository.CreateProcess(imageEntity, imageProcesses)
}
