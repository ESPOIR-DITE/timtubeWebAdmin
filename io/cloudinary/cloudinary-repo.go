package cloudinary

import (
	"context"
	"fmt"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func GetCld() (*cloudinary.Cloudinary, context.Context) {
	cld, _ := cloudinary.NewFromParams("<your-cloud-name>", "<your-api-key>", "<your-api-secret>")
	ctx := context.Background()
	return cld, ctx
}

func UploadVideo(file string) (*uploader.UploadResult, error) {
	cld, ctx := GetCld()
	resp, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{})
	if err != nil {
		fmt.Println(err, " error uploading video")
	}
	return resp, err
}
