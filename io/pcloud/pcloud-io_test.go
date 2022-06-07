package pcloud

import (
	"fmt"
	"testing"
)

func TestCreateFolder(t *testing.T) {
	CreateFolder()
}
func TestInit(t *testing.T) {
	Init()
}
func TestReadFiles(t *testing.T) {
	fmt.Println(ReadFiles())
}
