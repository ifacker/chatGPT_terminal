package file

import (
	"os"
)

// 写入文件
func WriteFile(filepath string, content []byte) (bool, error) {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return false, err
	}
	file.Write(content)
	return true, nil
}
