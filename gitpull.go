package goutils

import (
	"os"
	"os/exec"
	"path/filepath"
	"peterSZW/goutils/logger"
	//	"time"
)

func GitPull() {

	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		logger.Error(err)
	}
	path, err := filepath.Abs(file)
	if err != nil {
		logger.Error(err)
	}
	exePath := filepath.Dir(path)
	//logger.Debug("****git sync exepath:", exePath, "*****")
	if false {
		logger.Trace("----git pull----", exePath)
	}

	cmd := exec.Command("git", "pull")
	//cmd.Path = exePath
	out, err2 := cmd.CombinedOutput()
	if err2 != nil {
		logger.Trace(string(out))
		logger.Error(err2)
		logger.Trace(err2.Error())
		logger.Info("git pull fail")
	} else {
		//logger.Trace("git pull success")
	}

	cmd = exec.Command("./FileEncryptor", "dec", "songxxyy")
	//cmd.Path = exePath
	out, err2 = cmd.CombinedOutput()
	if err2 != nil {
		logger.Trace(string(out))
		logger.Error(err2)
		logger.Trace(err2.Error())
		logger.Info("FileEncryptor fail")
	} else {
		logger.Trace(string(out))
	}

	//time.Sleep(time.Second * 10)

}
