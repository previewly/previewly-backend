package processor

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"wsw/backend/domain/image/path"
	"wsw/backend/domain/image/process/options"
	"wsw/backend/ent/types"

	"github.com/h2non/bimg"
)

type (
	resizeProcessor struct {
		pathProvider  path.PathProvider
		pathGenerator path.FilenameGenerator
		width         *int
		height        *int
	}
)

// GetHash implements Processor.
func (r resizeProcessor) GetHash() string { return r.GeneratePathPrefix() }

// Run implements Process.
func (r resizeProcessor) Run(from path.PathData, filename string) (*path.PathData, error) {
	to := r.generatePath(from, filename)
	buffer, err := bimg.Read(from.FullPath)
	if err != nil {
		return nil, err
	}

	size, err := bimg.NewImage(buffer).Size()
	if err != nil {
		return nil, err
	}
	ratio := float64(size.Width) / float64(size.Height)

	resizeWidth := r.getResizeWidth(ratio)
	resizeHeight := r.getResizeHeight(ratio)

	newImage, err := bimg.NewImage(buffer).Resize(resizeWidth, resizeHeight)
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

func (r resizeProcessor) generatePath(from path.PathData, filename string) *path.PathData {
	prefix := r.GeneratePathPrefix()
	newPath := r.pathGenerator.GenerateFilepath(&prefix)
	return r.pathProvider.Provide(newPath, filename)
}

func (r resizeProcessor) getResizeHeight(ratio float64) int {
	if r.height != nil {
		return *r.height
	}
	return int(ratio * float64(*r.width))
}

func (r resizeProcessor) getResizeWidth(ratio float64) int {
	if r.width != nil {
		return *r.width
	}
	return int(ratio * float64(*r.height))
}

func (r resizeProcessor) GeneratePathPrefix() string {
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

// newResizeProcessor implements ProcessFactory.
func newResizeProcessor(pathProvider path.PathProvider, pathGenerator path.FilenameGenerator, opts []types.ImageProcessOption) (Processor, error) {
	width := options.ExtractIntOption(opts, "width")
	height := options.ExtractIntOption(opts, "height")

	if width != nil || height != nil {
		return resizeProcessor{width: width, height: height, pathProvider: pathProvider, pathGenerator: pathGenerator}, nil
	}
	return nil, errors.New("width or height should be provided")
}
