package cloudinary

import (
	"context"

	cld "github.com/cloudinary/cloudinary-go/v2"
)

var client *cld.Cloudinary
var ctx context.Context

func initiate() error {
	var err error

	ctx = context.Background()

	client, err = cld.New()

	return err
}
