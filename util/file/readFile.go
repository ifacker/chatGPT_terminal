package file

import (
	"io"
	"os"
)

// 读取文件
func ReadFile(filePath string) ([]byte, error) {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return body, nil
}
