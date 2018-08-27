package goutils

import (
	"peterSZW/goutils/logger"
	//"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)

func CompressVoiceToMp3(filename1 string, filename2 string) {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	exePath := filepath.Dir(path)

	logger.Debug("****CompressVoiceToMp3 exepath:", exePath, "*****")
	cmd := exec.Command(exePath+"\\utils\\lame.exe", "--preset", "cbr", "18", exePath+filename1, exePath+filename2)
	cmd.Run()

	//return
	//stdout, err := cmd.StdoutPipe()
	//if err != nil {
	//	fmt.Println("StdoutPipe: " + err.Error())
	//	return
	//}

	//stderr, err := cmd.StderrPipe()
	//if err != nil {
	//	fmt.Println("StderrPipe: ", err.Error())
	//	return
	//}

	//if err := cmd.Start(); err != nil {
	//	fmt.Println("Start: ", err.Error())
	//	return
	//}

	//bytesErr, err := ioutil.ReadAll(stderr)
	//if err != nil {
	//	fmt.Println("ReadAll stderr: ", err.Error())
	//	return
	//}

	//if len(bytesErr) != 0 {
	//	fmt.Printf("stderr is not nil: %s", bytesErr)
	//	return
	//}

	//bytes, err := ioutil.ReadAll(stdout)
	//if err != nil {
	//	fmt.Println("ReadAll stdout: ", err.Error())
	//	return
	//}

	//if err := cmd.Wait(); err != nil {
	//	fmt.Println("Wait: ", err.Error())
	//	return
	//}

	//fmt.Printf("stdout: %s", bytes)
}
