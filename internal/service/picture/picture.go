package picture

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
	"github.com/pariip/notes-go/pkg/random"
	"io"
	"mime/multipart"
	"os"
)

func (s service) UploadImage(pic *multipart.FileHeader) (*params.ImageResponse, error) {
	//Open
	src, err := pic.Open()
	if err != nil {
		return nil, err
	}
	defer src.Close()

	//Create
	imageName := random.String(8) + pic.Filename
	dst, err := os.Create(s.cfg.ImagePath + imageName)
	if err != nil {
		return nil, err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return nil, err
	}
	image := &models.Picture{
		Name: pic.Filename,
		Alt:  s.cfg.LinkPath + imageName,
	}

	//Upload
	if image, err = s.imageRepo.UploadImage(image); err != nil {
		return nil, err
	}

	return &params.ImageResponse{
		ID:  image.ID,
		Url: image.Alt,
	}, nil

}
