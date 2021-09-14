package contract

import (
	"github.com/pariip/notes-go/internal/params"
	"mime/multipart"
)

type ImageService interface {
	UploadImage(pic *multipart.FileHeader) (*params.ImageResponse, error)
}
