package main

import (
	"strings"
)

func GetType(path string) string {
	s := strings.Split(path, "/")
	if len(s) == 0 {
		return "undefined"
	}
	file := s[len(s)-1]
	f := strings.Split(file, ".")
	if len(f) == 0 {
		return "undefined"
	}
	extension := f[len(f)-1]
	return extension
}
func main() {
	GetType("web/registration/app.js")
}
