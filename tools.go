package utils

import (
	"encoding/json"
	"strings"
)

func GetFileExtension(Filename string) string {
	Result := ""
	f := strings.Split(Filename, ".")
	Result = f[len(f)-1]
	return Result
}
func ToJsonString(a interface{}) string {
	b, err := json.Marshal(a)
	if err != nil {
		return ``
	} else {
		return string(b)
	}
}
