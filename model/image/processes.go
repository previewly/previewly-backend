package image

import (
	"wsw/backend/ent"
	"wsw/backend/ent/repository"
	"wsw/backend/ent/types"

	"github.com/xorcare/pointer"
)

type (
	ImageProcesses interface {
		Create(entity *ent.Image, processes []types.ImageProcess, hash string) (*ent.ImageProcess, error)
		Update(*ent.ImageProcess, string, types.StatusEnum, string) (*ent.ImageProcess, error)
		TryGetByHash(imageID int, processesHash string) (*ent.ImageProcess, error)
	}
	imageProcessesImpl struct {
		processRepository repository.ImageProcessRepository
		imageRepository   repository.ImageRepository
	}
)

func (i imageProcessesImpl) TryGetByHash(imageID int, processesHash string) (*ent.ImageProcess, error) {
	entity, err := i.processRepository.GetByHash(imageID, processesHash)
	if ent.IsNotFound(err) {
		return nil, nil
	}
	return entity, err
}

func (i imageProcessesImpl) Update(processEntity *ent.ImageProcess, prefix string, status types.StatusEnum, err string) (*ent.ImageProcess, error) {
	return i.processRepository.Update(processEntity, prefix, status, pointer.String(err))
}

func NewImageProcesses(processRepository repository.ImageProcessRepository, imageRepository repository.ImageRepository) ImageProcesses {
	return imageProcessesImpl{processRepository: processRepository, imageRepository: imageRepository}
}

func (i imageProcessesImpl) Create(entity *ent.Image, processes []types.ImageProcess, hash string) (*ent.ImageProcess, error) {
	return i.imageRepository.CreateProcess(entity, processes, hash)
}
