package file

import (
	"fmt"
	"testing"
)

func TestReadFile(t *testing.T) {
	body, err := ReadFile("test.txt")
	fmt.Println(string(body), err)
}
