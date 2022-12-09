package ufile

import (
	uuid "github.com/satori/go.uuid"
	"net/http"
	"os"
)

func ExtFile(path string) string {
	for i := len(path) - 1; i >= 0 && path[i] != '/'; i-- {
		if path[i] == '.' {
			return path[i:]
		}
	}
	return ""
}

func GenFileName() string {
	return "/"+ uuid.NewV4().String()
}

func GetFileContentType(out *os.File) (string, error) {

	// 只需要前 512 个字节就可以了
	buffer := make([]byte, 512)

	_, err := out.Read(buffer)
	if err != nil {
		return "", err
	}

	contentType := http.DetectContentType(buffer)

	return contentType, nil
}
