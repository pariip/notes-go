package postgres

import (
	"github.com/pariip/notes-go/internal/db/postgres/schema"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

func (r *repository) CreateToken(token string, userID uint) error {
	t := &schema.Token{
		Value:  token,
		UserID: userID,
	}
	if err := r.db.Model(&schema.Token{}).Where("user_id =?", userID).Delete(&schema.Token{}).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.auth",
			Function: "CreateToken",
			Params:   map[string]interface{}{"user_id": userID},
			Message:  err.Error(),
		})
		return cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	if err := r.db.Model(&schema.Token{}).Create(t).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.auth",
			Function: "CreateToken",
			Params:   map[string]interface{}{"user_id": userID},
			Message:  err.Error(),
		})
		return cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	return nil
}

func (r *repository) TokenIsExistWithUserID(token string, userID uint) (bool, error) {
	t := &schema.Token{}
	if err := r.db.Model(&schema.Token{}).Where("value = ? and user_id = ?", token, userID).First(&t).Error; err != nil {

		if isErrorNotFound(err) {
			return false, err
		}
		r.logger.Error(&log.Field{
			Section:  "repository.auth",
			Function: "TokenIsExistWithUserID",
			Message:  err.Error(),
		})

		return false, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	return true, nil
}
