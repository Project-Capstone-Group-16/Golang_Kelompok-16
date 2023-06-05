package usecase

import (
	"Capstone/utils"
	"context"
	"os"

	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadImageCloud(image string, ctx context.Context) (resp *uploader.UploadResult, err error) {
	file, err := os.Open(image)
	cld := utils.Credentials()
	resp, err = cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       "inventron",
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true)})
	if err != nil {
		return
	}

	return resp, nil
}
