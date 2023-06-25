package utils

import (
	"Capstone/models/payload"
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

func Credentials() *cloudinary.Cloudinary {
	cld, _ := cloudinary.New()
	cld.Config.URL.Secure = true
	return cld
}

func UploadImageCloud(fileHeader *multipart.FileHeader) (imageUrl string, err error) {

	file, _ := fileHeader.Open()

	cld := Credentials()

	resp, err := cld.Upload.Upload(context.Background(), file, uploader.UploadParams{
		PublicID:       "Inventron" + "/" + fileHeader.Filename,
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})

	if err != nil {
		return
	}

	imageUrl = resp.SecureURL

	return imageUrl, nil
}

func UploadImageCloudBase64(req *payload.UploadImageCloudinaryBase64) (imageUrl string, err error) {
	cld := Credentials()
	uuid := uuid.New()

	resp, err := cld.Upload.Upload(context.Background(), req.Image, uploader.UploadParams{
		PublicID:       "Inventron/" + uuid.String(),
		UniqueFilename: api.Bool(false),
		Overwrite:      api.Bool(true),
	})

	if err != nil {
		return "", err
	}

	imageUrl = resp.SecureURL

	return imageUrl, nil
}
