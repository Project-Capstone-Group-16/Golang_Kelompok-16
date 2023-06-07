package utils

import (
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func Credentials() *cloudinary.Cloudinary {
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	return cld
}

func UploadImageCloud(image string, ctx context.Context) (resp *uploader.UploadResult, err error) {
	file, err := os.Open(image)

	cld := Credentials()

	resp, err = cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       "inventron",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})

	if err != nil {
		return
	}

	return resp, nil
}
