package input

import (
	"wsw/backend/ent"
	"wsw/backend/ent/types"
)

type (
	Input struct {
		Image     *ent.Image
		Processes []types.ImageProcess
	}
)
