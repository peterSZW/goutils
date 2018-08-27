package utils

import (
	"math/rand"
	"time"
)

func CreateRandomNumber(endNum int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(endNum)
}
