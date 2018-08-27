package utils

import (
	"fmt"
	"strings"
)

func GenNextMemberCode(code int) int {
	r := code + 1
	s := fmt.Sprint(r)

	for strings.ContainsAny(s, "47") {
		r = r + 1
		s = fmt.Sprint(r)
	}
	return r
}
