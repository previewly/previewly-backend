package processor

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"wsw/backend/domain/image/path"
	"wsw/backend/domain/image/process/options"
	processPathProvider "wsw/backend/domain/image/process/path"
	"wsw/backend/domain/image/process/processor/sizes"
	"wsw/backend/ent/types"
	"wsw/backend/lib/utils"

	"github.com/h2non/bimg"
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
	bimgNewImage := bimg.NewImage(buffer)

	size, err := bimgNewImage.Size()
	if err != nil {
		return nil, err
	}

	newSizes := sizes.GetNewSizesByRatio(size, c.width, c.height)
	cropWidth := newSizes.Width
	cropHeight := newSizes.Height
	x := (size.Width - cropWidth) / 2
	y := (size.Height - cropHeight) / 2

	utils.D(x, y, cropWidth, cropHeight, size)

	newImage, err := bimgNewImage.Extract(y, x, cropWidth, cropHeight)
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
