package postgres

import (
	"github.com/pariip/notes-go/internal/db/postgres/schema"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

//UploadImage save a image in the database
func (r *repository) UploadImage(pic *models.Picture) (*models.Picture, error) {
	p := schema.ConvertPicture(pic)

	if err := r.db.Create(p).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.picture",
			Function: "UploadImage",
			Params:   map[string]interface{}{"picture": p},
			Message:  err.Error(),
		})
		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	return p.ConvertModel(), nil
}

//IsImageExist check image is exist
func (r *repository) IsImageExist(picAlt string) (*models.Picture, error) {
	picture := &schema.Picture{}
	if err := r.db.Model(&schema.Picture{}).Where("alt = ?", picAlt).First(picture).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "IsImageExist",
			Params:   map[string]interface{}{"picture": picture.Name},
			Message:  err.Error(),
		})
		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	return picture.ConvertModel(), nil
}
