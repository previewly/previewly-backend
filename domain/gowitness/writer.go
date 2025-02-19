package gowitness

import (
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"

	"wsw/backend/domain/image"
	"wsw/backend/domain/path/screenshot/relative"
	"wsw/backend/ent"
	"wsw/backend/ent/repository"

	"github.com/sensepost/gowitness/pkg/models"
	"github.com/sensepost/gowitness/pkg/writers"
	"github.com/xorcare/pointer"
)

type (
	Writer interface {
		writers.Writer
		Error(error) error
	}
	writerImpl struct {
		url                  *ent.Url
		urlRepository        repository.Url
		statRepository       repository.Stat
		relativePathProvider relative.Provider
		saver                image.Saver
	}
)

// Error implements Writer.
func (w writerImpl) Error(err error) error {
	_, saveErr := w.urlRepository.SaveFailure(err.Error(), w.url.ID)
	return saveErr
}

// LeftTrucate a string if its more than max
func (w writerImpl) leftTrucate(s string, max int) string {
	if len(s) <= max {
		return s
	}

	return s[max:]
}

// SafeFileName takes a string and returns a string safe to use as
// a file name.
func (w writerImpl) safeFileName(s string) string {
	var builder strings.Builder

	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '.' {
			builder.WriteRune(r)
		} else {
			builder.WriteRune('-')
		}
	}

	return builder.String()
}

// Write implements writers.Writer.
func (w writerImpl) Write(result *models.Result) error {
	filename := w.newFilename(result.URL)

	imageEntity, err := w.addScreenshotToImages(result.Screenshot, filename)
	if err != nil {
		return err
	}

	statEntity, errStat := w.statRepository.Insert(&result.Title, imageEntity)
	if errStat != nil {
		return errStat
	}

	_, errStatus := w.urlRepository.SaveSuccess(w.relativePathProvider.Provide(filename), statEntity, w.url.ID)
	if errStatus != nil {
		return errStatus
	}

	return nil
}

func (w writerImpl) newFilename(url string) string {
	return w.leftTrucate(w.safeFileName(url)+".jpeg", 200)
}

func (w writerImpl) addScreenshotToImages(screenshotContent string, filename string) (*ent.Image, error) {
	reader, err := w.decodeScreenReader(screenshotContent)
	if err != nil {
		return nil, err
	}

	imageEntity, errImage := w.saver.SaveImage(filename, reader, "image/jpeg", w.newExtra(w.url.ID))
	if errImage != nil {
		return nil, errImage
	}
	return imageEntity, nil
}

func (w writerImpl) newExtra(id int) *string {
	var builder strings.Builder

	builder.WriteString("url: ")
	builder.WriteString(strconv.Itoa(id))
	return pointer.String(builder.String())
}

func (w writerImpl) decodeScreenReader(screenshot string) (io.ReadSeeker, error) {
	content, err := base64.StdEncoding.DecodeString(screenshot)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 string: %v", err)
	}
	return strings.NewReader(string(content)), nil
}

func NewRunnerWriter(
	url *ent.Url,
	urlRepository repository.Url,
	statRepository repository.Stat,
	relativePathProvider relative.Provider,
	saver image.Saver,
) Writer {
	return writerImpl{
		url:                  url,
		urlRepository:        urlRepository,
		statRepository:       statRepository,
		relativePathProvider: relativePathProvider,
		saver:                saver,
	}
}
