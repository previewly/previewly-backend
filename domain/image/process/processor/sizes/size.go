package sizes

import (
	"github.com/h2non/bimg"
)

func GetRatio(size bimg.ImageSize) float32 { return float32(size.Width) / float32(size.Height) }

func GetNewSizesByRatio(imageSize bimg.ImageSize, toWidth *int, toHeight *int) bimg.ImageSize {
	if toWidth != nil && toHeight != nil {
		return bimg.ImageSize{Width: *toWidth, Height: *toHeight}
	}

	ratio := GetRatio(imageSize)

	if toWidth != nil {
		return bimg.ImageSize{Width: *toWidth, Height: int(float32(*toWidth) / ratio)}
	}
	if toHeight != nil {
		return bimg.ImageSize{Width: int(float32(*toHeight) * ratio), Height: *toHeight}
	}

	return bimg.ImageSize{Width: 0, Height: 0}
}
