package util

import "os"

func CloudinaryPreset() string {
	return os.Getenv("CLOUDINARY_PRESET")
}
