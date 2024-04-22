package cloudinary

import (
	"goRepositoryPattern/util"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFile(file string) (*uploader.UploadResult, error) {
	err := initiate()
	if err != nil {
		return nil, err
	}

	res, err := client.Upload.Upload(ctx, file, uploader.UploadParams{
		UploadPreset: util.CloudinaryPreset(),
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}
