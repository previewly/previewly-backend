package image

import (
	"context"

	"wsw/backend/ent"
	"wsw/backend/ent/repository"
	"wsw/backend/ent/types"
)

type (
	ImageProcesses interface {
		Create(*ent.UploadImage, []types.ImageProcess) (*ent.ImageProcess, error)
		Save(context.Context, *ent.ImageProcess) (*ent.ImageProcess, error)
	}
	imageProcessesImpl struct {
		processRepository repository.ImageProcessRepository
		imageRepository   repository.UploadImageRepository
	}
)

func NewImageProcesses(processRepository repository.ImageProcessRepository, imageRepository repository.UploadImageRepository) ImageProcesses {
	return imageProcessesImpl{processRepository: processRepository, imageRepository: imageRepository}
}

func (i imageProcessesImpl) Save(ctx context.Context, processEntity *ent.ImageProcess) (*ent.ImageProcess, error) {
	panic("unimplemented")
}

func (i imageProcessesImpl) Create(imageEntity *ent.UploadImage, imageProcesses []types.ImageProcess) (*ent.ImageProcess, error) {
	return i.imageRepository.CreateProcess(imageEntity, imageProcesses)
}
