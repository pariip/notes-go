package postgres

import (
	"github.com/pariip/notes-go/internal/db/postgres/schema"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

func (r *repository) CreateUser(user *models.User) (*models.User, error) {
	u := schema.ConvertUser(user)
	if err := r.db.Create(u).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "CreateUser",
			Params:   map[string]interface{}{"user": u},
			Message:  err.Error(),
		})
		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	return u.ConvertModel(), nil

}

func (r *repository) GetUserByID(userID uint) (*models.User, error) {
	user := new(schema.User)
	if err := r.db.Model(&schema.User{}).Where("id =?", userID).First(user).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "GetUserByID",
			Params:   map[string]interface{}{"userID": userID},
			Message:  err.Error(),
		})
		if isErrorNotFound(err) {
			return nil, cerrors.New(cerrors.KindNotFound, messages.UserNotFound)
		}
		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)

	}
	return user.ConvertModel(), nil
}

func (r *repository) GetUserByUsername(username string) (*models.User, error) {
	user := new(schema.User)

	if err := r.db.Model(&schema.User{}).Where("username =?", username).First(user).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "GetUserByUsername",
			Params:   map[string]interface{}{"username": username},
			Message:  err.Error(),
		})
		if isErrorNotFound(err) {
			return nil, cerrors.New(cerrors.KindNotFound, messages.UserNotFound)
		}
		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)

	}
	return user.ConvertModel(), nil

}

func (r *repository) UpdateUser(user *models.User) (*models.User, error) {
	u := schema.ConvertUser(user)

	if err := r.db.Model(&schema.User{}).First(&schema.User{}, u.ID).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "UpdateUser",
			Params:   map[string]interface{}{"user": u},
			Message:  err.Error(),
		})

		if isErrorNotFound(err) {
			return nil, cerrors.New(cerrors.KindNotFound, messages.UserNotFound)
		}

		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	if err := r.db.Model(&schema.User{}).Where("id = ?", u.ID).Save(u).Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "UpdateUser",
			Params:   map[string]interface{}{"user": u},
			Message:  err.Error(),
		})

		return nil, cerrors.New(cerrors.KindUnexpected, messages.DBError)

	}
	return u.ConvertModel(), nil
}

func (r *repository) DeleteUser(user *models.User) error {
	u := schema.ConvertUser(user)
	res := r.db.Where("id = ?", u.ID).Delete(u)
	if err := res.Error; err != nil {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "DeleteUser",
			Params:   map[string]interface{}{"user": u},
			Message:  err.Error(),
		})

		return cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	if res.RowsAffected != 1 {
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "DeleteUser",
			Params:   map[string]interface{}{"user": u},
			Message:  r.translator.TranslateEn(messages.UserNotFound),
		})

		return cerrors.New(cerrors.KindNotFound, messages.UserNotFound)
	}
	return nil
}

func (r *repository) IsUsernameExist(username string) (bool, error) {
	user := &schema.User{}

	if err := r.db.Model(&schema.User{}).Where("username= ?", username).First(user).Error; err != nil {
		if isErrorNotFound(err) {
			return false, nil
		}
		r.logger.Error(&log.Field{
			Section:  "repository.user",
			Function: "IsUsernameExist",
			Params:   map[string]interface{}{"username": username},
			Message:  err.Error(),
		})

		return false, cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}

	return true, nil
}
