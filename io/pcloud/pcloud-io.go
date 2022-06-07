package pcloud

import (
	"fmt"
	"github.com/yanmhlv/pcloud"
	"os"
)

var c = pcloud.NewClient()

func Init() {
	if c.Auth == nil {
		c.Login("espoirditekemena@gmail.com", "lunchbar@pcloud")
	}
}
func CreateFolder() {
	Init()
	fmt.Println("Login", c.Login("espoirditekemena@gmail.com", "lunchbar@pcloud"))
	fmt.Println("CreateFolder /helloworld", c.CreateFolder("/timtubevideo", 0, ""))
}
func UploadFile(file *os.File) {
	Init()
	c.UploadFile(file, "", 0, "index.html", 0, "", 0)
}
func ReadFiles() []string {
	Init()
	result, err := c.GetFileLink(0, "timtubevideo/", 1, "application/pdf", 1, 1)
	if err != nil {
		fmt.Println(err, " error reading file")
	}
	return result
}
