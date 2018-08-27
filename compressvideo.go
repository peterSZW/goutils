package utils

import (
	"peterSZW/goutils/logger"
	//"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func CompressMp4(filename1 string, filename2 string) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		logger.Error(err)
	}
	path, err := filepath.Abs(file)
	if err != nil {
		logger.Error(err)
	}
	exePath := filepath.Dir(path)
	logger.Debug("****CompressVoiceToMp3 exepath:", exePath, "*****")
	cmd := exec.Command(exePath+"\\utils\\ffmpeg.exe", "-i", exePath+filename1, "-acodec", "aac", "-ac", "1", "-ar", "18k", "-s", "320x240", "-b", "120000", "-r", "25", "-vcodec", "libx264", exePath+filename2)
	err2 := cmd.Run()
	if err2 != nil {
		logger.Error(err2)
		logger.Info("压缩视频失败")
	} else {
		logger.Info("压缩视频成功")
	}
}
