package processor

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"wsw/backend/domain/image/path"
	"wsw/backend/domain/image/process/options"
	"wsw/backend/domain/image/process/processor/sizes"

	processPathProvider "wsw/backend/domain/image/process/path"
	"wsw/backend/ent/types"

	"github.com/h2non/bimg"
	"github.com/xorcare/pointer"
)

type (
	resizeProcessor struct {
		pathProvider processPathProvider.Provider
		width        *int
		height       *int
	}
)

// GetHash implements Processor.
func (r resizeProcessor) GetHash() string { return r.getPathPrefix() }

// Run implements Process.
func (r resizeProcessor) Run(from path.PathData, filename string) (*path.PathData, error) {
	to := r.pathProvider.Get(r.getPathPrefix(), filename)

	buffer, err := bimg.Read(from.FullPath)
	if err != nil {
		return nil, err
	}
	bimgNewImage := bimg.NewImage(buffer)

	size, err := bimgNewImage.Size()
	if err != nil {
		return nil, err
	}

	newSizes := sizes.GetNewSizesByRatio(size, r.width, r.height)

	newImage, err := bimgNewImage.Resize(newSizes.Width, newSizes.Height)
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

func (r resizeProcessor) getRatio(buffer []byte) (*float32, error) {
	size, err := bimg.NewImage(buffer).Size()
	if err != nil {
		return nil, err
	}
	return pointer.Float32(float32(size.Width) / float32(size.Height)), nil
}

func (r resizeProcessor) getResizeHeight(ratio float32) int {
	if r.height != nil {
		return *r.height
	}
	return int(ratio * float32(*r.width))
}

func (r resizeProcessor) getResizeWidth(ratio float32) int {
	if r.width != nil {
		return *r.width
	}
	return int(ratio * float32(*r.height))
}

func (r resizeProcessor) getPathPrefix() string {
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

// NewResizeProcessor implements ProcessFactory.
func NewResizeProcessor(pathProvider processPathProvider.Provider, opts []types.ImageProcessOption) (Processor, error) {
	width := options.ExtractIntOption(opts, "width")
	height := options.ExtractIntOption(opts, "height")

	if width != nil || height != nil {
		return resizeProcessor{width: width, height: height, pathProvider: pathProvider}, nil
	}
	return nil, errors.New("width or height should be provided")
}
