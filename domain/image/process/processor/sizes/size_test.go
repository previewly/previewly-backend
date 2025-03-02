package sizes_test

import (
	"testing"

	"wsw/backend/domain/image/process/processor/sizes"

	"github.com/stretchr/testify/assert"
	"github.com/xorcare/pointer"

	"github.com/h2non/bimg"
)

type (
	ratioTest struct {
		size     bimg.ImageSize
		expected float32
	}

	newSizeTest struct {
		imageSize bimg.ImageSize
		toWidth   *int
		toHeight  *int
		expected  bimg.ImageSize
	}
)

var (
	imageSizeRatio1    = bimg.ImageSize{Width: 100, Height: 100}
	imageSizeRatio2    = bimg.ImageSize{Width: 200, Height: 100}
	imageSizeRatioHalf = bimg.ImageSize{Width: 100, Height: 200}

	ratioTests = []ratioTest{
		{size: imageSizeRatioHalf, expected: 0.5},
		{size: imageSizeRatio2, expected: 2},
		{size: imageSizeRatio1, expected: 1},
	}

	sizeTests = []newSizeTest{
		{imageSize: imageSizeRatio1, toWidth: pointer.Int(10), toHeight: nil, expected: bimg.ImageSize{Width: 10, Height: 10}},
		{imageSize: imageSizeRatio1, toWidth: nil, toHeight: pointer.Int(10), expected: bimg.ImageSize{Width: 10, Height: 10}},

		{imageSize: imageSizeRatio2, toWidth: pointer.Int(10), toHeight: nil, expected: bimg.ImageSize{Width: 10, Height: 5}},
		{imageSize: imageSizeRatio2, toWidth: nil, toHeight: pointer.Int(10), expected: bimg.ImageSize{Width: 20, Height: 10}},

		{imageSize: imageSizeRatioHalf, toWidth: pointer.Int(10), toHeight: nil, expected: bimg.ImageSize{Width: 10, Height: 20}},
		{imageSize: imageSizeRatioHalf, toWidth: nil, toHeight: pointer.Int(10), expected: bimg.ImageSize{Width: 5, Height: 10}},
	}
)

func TestRatio(t *testing.T) {
	for _, test := range ratioTests {
		if output := sizes.GetRatio(test.size); output != test.expected {
			t.Errorf("Output %f not equal to expected %f", output, test.expected)
		}
	}
}

func TestSizes(t *testing.T) {
	assert := assert.New(t)
	for _, test := range sizeTests {
		got := sizes.GetNewSizesByRatio(test.imageSize, test.toWidth, test.toHeight)
		assert.Equal(test.expected.Width, got.Width)
		assert.Equal(test.expected.Height, got.Height)
	}
}
