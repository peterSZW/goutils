package utils

import (
	"image/jpeg"
	"io"
	"net/http"
	"os"
	"peterSZW/goutils/logger"

	"github.com/nfnt/resize"
)

func DownloadJpg(url string, filename string) bool {
	resp, err := http.Get(url)
	if err != nil {
		logger.Error(err)
		return false
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		logger.Error(err)
	}
	defer out.Close()

	n, err := io.Copy(out, resp.Body)
	if err != nil {
		logger.Error(err)
		return false
	} else {
		logger.Debug("FILEsize", n)
		return true
	}

}
func FileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		logger.Error(err)
	}
	return err == nil || os.IsExist(err)
}

func CopyFile(src, dst string) (w int64, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		logger.Error(err)
		return
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		logger.Error(err)
		return
	}

	defer dstFile.Close()

	return io.Copy(dstFile, srcFile)
}

func DownloadAndCompress(url, filename, filename2 string) {
	logger.Debug("DownloadAndCompress", url, filename)
	DownloadJpg(url, filename)
	if !FileExist(filename) {
		CopyFile("defaultFace.jpg", filename)
	}
	CompressJpg(filename, filename2)
}

func CompressJpg(filename string, filename2 string) {
	file, err := os.Open(filename)
	if err != nil {
		logger.Error(err)
		return
	}
	defer file.Close()

	// decode jpeg into image.Image
	//	png1, err := png.Decode(file)

	img, err := jpeg.Decode(file)
	if err != nil {
		logger.Error(err)
		return
	}

	// resize to width 1000 using Lanczos resampling
	// and preserve aspect ratio
	m := resize.Resize(32, 0, img, resize.Lanczos3)

	out, err := os.Create(filename2)
	if err != nil {
		logger.Error(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
