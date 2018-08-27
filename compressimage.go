package goutils

import (
	"image"
	"image/jpeg"
	"os"
	"peterSZW/goutils/logger"

	"github.com/nfnt/resize"
)

func CompressImageToSmall(filename1 string, filename2 string) bool {
	file, err := os.Open(filename1)
	if err != nil {
		logger.Error(err)
	}

	// decode jpeg into image.Image
	img, _, err := image.Decode(file)
	if err != nil {
		logger.Error(err)
	}
	file.Close()

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(100, 0, img, resize.Lanczos3)

	out, err := os.Create(filename2)
	if err != nil {
		logger.Error(err)
	}
	defer out.Close()

	// write new image to file

	jpeg.Encode(out, m, nil)

	return true
}

func CompressImageToMiddle(filename1 string, filename2 string) bool {
	file, err := os.Open(filename1)
	if err != nil {
		logger.Error(err)
	}

	// decode jpeg into image.Image
	img, _, err := image.Decode(file)
	if err != nil {
		logger.Error(err)
	}
	file.Close()
	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(500, 0, img, resize.Lanczos3)
	out, err := os.Create(filename2)
	if err != nil {
		logger.Error(err)
	}
	defer out.Close()

	// write new image to file

	jpeg.Encode(out, m, nil)

	return true
}
