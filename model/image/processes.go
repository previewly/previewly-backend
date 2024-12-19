package image

import (
	"context"

	"wsw/backend/ent"
	"wsw/backend/ent/types"
)

type (
	ImageProcesses interface {
		Create(*ent.UploadImage, []*types.ImageProcess) (*ent.ImageProcess, error)
		Save(context.Context, *ent.ImageProcess) (*ent.ImageProcess, error)
	}
	imageProcessesImpl struct{}
)

func (i imageProcessesImpl) Save(ctx context.Context, processEntity *ent.ImageProcess) (*ent.ImageProcess, error) {
	panic("unimplemented")
}

func (i imageProcessesImpl) Create(imageEntity *ent.UploadImage, imageProcesses []*types.ImageProcess) (*ent.ImageProcess, error) {
	panic("unimplemented")
}

func NewImageProcesses() ImageProcesses {
	return imageProcessesImpl{}
}
