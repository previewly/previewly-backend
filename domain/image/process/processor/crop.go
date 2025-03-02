package processor

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"wsw/backend/domain/image/path"
	"wsw/backend/domain/image/process/options"
	processPathProvider "wsw/backend/domain/image/process/path"
	"wsw/backend/ent/types"

	"github.com/h2non/bimg"
	"github.com/xorcare/pointer"
)

type (
	cropProcessor struct {
		pathProvider processPathProvider.Provider
		width        *int
		height       *int
	}
)

// GetHash implements Processor.
func (c cropProcessor) GetHash() string { return c.getPathPrefix() }

// Run implements Processor.
func (c cropProcessor) Run(from path.PathData, filename string) (*path.PathData, error) {
	to := c.pathProvider.Get(c.getPathPrefix(), filename)

	buffer, err := bimg.Read(from.FullPath)
	if err != nil {
		return nil, err
	}

	size, err := bimg.NewImage(buffer).Size()
	if err != nil {
		return nil, err
	}

	ratio := pointer.Float32(float32(size.Width) / float32(size.Height))

	cropWidth := c.getResizeWidth(*ratio)
	cropHeight := c.getResizeHeight(*ratio)

	x := (size.Width - cropWidth) / 2
	y := (size.Height - cropHeight) / 2

	newImage, err := bimg.NewImage(buffer).Extract(y, x, cropWidth, cropHeight)
	if err != nil {
		return nil, err
	}

	// Create the uploads folder if it doesn't already exist
	errMkdir := os.MkdirAll(to.Directory, os.ModePerm)
	if errMkdir != nil {
		return nil, errMkdir
	}

	errWrite := bimg.Write(to.FullPath, newImage)
	if errWrite != nil {
		return nil, errWrite
	}
	return to, nil
}

func NewCropProcessor(pathProvider processPathProvider.Provider, opts []types.ImageProcessOption) (Processor, error) {
	width := options.ExtractIntOption(opts, "width")
	height := options.ExtractIntOption(opts, "height")

	if width != nil || height != nil {
		return cropProcessor{width: width, height: height, pathProvider: pathProvider}, nil
	}
	return nil, errors.New("width or height should be provided")
}

func (c cropProcessor) getPathPrefix() string {
	var sb strings.Builder
	sb.WriteString("crop/")
	if c.width != nil {
		sb.WriteString(strconv.Itoa(*c.width))
		sb.WriteString("/")
	}
	if c.height != nil {
		sb.WriteString(strconv.Itoa(*c.height))
		sb.WriteString("/")
	}
	return sb.String()
}

func (c cropProcessor) getResizeHeight(ratio float32) int {
	if c.height != nil {
		return *c.height
	}
	return int(ratio * float32(*c.width))
}

func (c cropProcessor) getResizeWidth(ratio float32) int {
	if c.width != nil {
		return *c.width
	}
	return int(ratio * float32(*c.height))
}
