package postgres

import (
	"github.com/pariip/notes-go/internal/db/postgres/schema"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

func (r *repository) migration() error {
	err := r.db.AutoMigrate(&schema.User{}, &schema.Note{}, &schema.Picture{})
	if err != nil {
		r.logger.Error(&log.Field{
			Section:  "postgres.migration",
			Function: "migration",
			Message:  err.Error(),
		})
		return cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	return nil

}
