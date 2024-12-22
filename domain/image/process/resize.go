package process

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"wsw/backend/domain/image/path"
	"wsw/backend/ent/types"

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
	buffer, err := bimg.Read(from.FullPath)
	if err != nil {
		return err
	}

	size, err := bimg.NewImage(buffer).Size()
	if err != nil {
		return err
	}
	ratio := float64(size.Width) / float64(size.Height)

	resizeWidth := r.getResizeWidth(ratio)
	resizeHeight := r.getResizeHeight(ratio)

	newImage, err := bimg.NewImage(buffer).Resize(resizeWidth, resizeHeight)
	if err != nil {
		return err
	}

	// Create the uploads folder if it doesn't already exist
	errMkdir := os.MkdirAll(to.Directory, os.ModePerm)
	if errMkdir != nil {
		return errMkdir
	}

	return bimg.Write(to.FullPath, newImage)
}

func (r resizeProcessImpl) getResizeHeight(ratio float64) int {
	if r.height != nil {
		return *r.height
	}
	return int(ratio * float64(*r.width))
}

func (r resizeProcessImpl) getResizeWidth(ratio float64) int {
	if r.width != nil {
		return *r.width
	}
	return int(ratio * float64(*r.height))
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
