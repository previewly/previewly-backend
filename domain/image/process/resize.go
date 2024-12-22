package process

import (
	"errors"
	"strconv"
	"strings"

	"wsw/backend/domain/image/path"
	"wsw/backend/ent/types"
	"wsw/backend/lib/utils"

	"github.com/h2non/bimg"
)

type (
	resizeProcessFactoryImpl struct{}
	resizeProcessImpl        struct {
		width  *int
		height *int
	}
)

// Run implements Process.
func (r resizeProcessImpl) Run(from path.PathData, to path.PathData) error {
	_, err := bimg.Read(from.FullPath)
	if err != nil {
		return err
	}
	utils.D(to)

	return nil
}

func (r resizeProcessImpl) GeneratePathPrefix() string {
	var sb strings.Builder
	sb.WriteString("resize/")
	if r.width != nil {
		sb.WriteString(strconv.Itoa(*r.width))
		sb.WriteString("/")
	}
	if r.height != nil {
		sb.WriteString(strconv.Itoa(*r.height))
		sb.WriteString("/")
	}
	return sb.String()
}

// Create implements ProcessFactory.
func (r resizeProcessFactoryImpl) Create(options []types.ImageProcessOption) (Process, error) {
	width := GetIntOption(options, "width")
	height := GetIntOption(options, "height")

	if width != nil || height != nil {
		return resizeProcessImpl{width: width, height: height}, nil
	}
	return nil, errors.New("width or height should be provided")
}
