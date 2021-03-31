package crypto

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
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
func CheckAccessArea(path string) (string, bool) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Println(err)
		return "", false
	}
	compare := strings.Contains(abs, dir)
	//fmt.Println("PWD:", dir)
	//fmt.Println("Absolute:", abs)
	ex := fileExists(abs)
	//fmt.Printf("Contains: %t | FileExists: %t\n", compare, ex)
	if !compare || !ex {
		return "", false
	}
	return abs, true
}

//func main() {
//	fmt.Println(CheckAccessArea("/frame/t"))

//}
