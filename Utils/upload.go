package utils

import (
	"context"
	"log"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(profilePictureFile *multipart.FileHeader) (string, error) {
	cld, err := cloudinary.New()
	if err != nil {
		log.Fatalf("Failed to intialize Cloudinary, %v", err)
	}
	var ctx = context.Background()
	uploadResult, err := cld.Upload.Upload(
		ctx, profilePictureFile,
		uploader.UploadParams{PublicID: "models",
			UniqueFilename: api.Bool(false),
			Overwrite:      api.Bool(true)})
	if err != nil {
		log.Fatalf("Failed to upload file, %v\n", err)
	}
	return uploadResult.SecureURL, err
}
