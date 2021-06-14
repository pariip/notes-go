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
