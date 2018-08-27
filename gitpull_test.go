package utils

import (
	"testing"
)

func init() {

}

func TestGitPull(t *testing.T) {
	//GitPull()
}

func TestMD5(t *testing.T) {
	if Md5string("hello") != "5d41402abc4b2a76b9719d911017c592" {
		t.Fail()
	}
}
