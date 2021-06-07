package user

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/hash"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

func (s *service) CreateUser(req *params.CreateUserRequest) (*models.User, error) {

	if err := s.validate.Username(req.Username); err != nil {
		return nil, err
	}

	if err := s.validate.Password(req.Password); err != nil {
		return nil, err
	}

	ok, err := s.userRepo.IsUsernameExist(req.Username)
	if err != nil {
		return nil, err
	}

	if ok {
		s.logger.Error(&log.Field{
			Section:  "service.user",
			Function: "CreateUser",
			Params:   map[string]interface{}{"req": req},
			Message:  s.translator.TranslateEn(messages.UsernameIsDuplicate),
		})
		return nil, cerrors.New(cerrors.KindInvalid, messages.UsernameIsDuplicate)

	}
	password, err := hash.Password(req.Password)
	if err != nil {
		s.logger.Error(&log.Field{
			Section:  "service.user",
			Function: "CreateUser",
			Params:   map[string]interface{}{"password": req.Password},
			Message:  err.Error(),
		})

		return nil, cerrors.New(cerrors.KindUnexpected, messages.GeneralError)
	}

	user := &models.User{
		Username:              req.Username,
		Password:              password,
		FirstName:             req.FirstName,
		LastName:              req.LastName,
		Email:                 req.Email,
		IsEmailVerified:       req.IsEmailVerified,
		PhoneNumber:           req.PhoneNumber,
		IsPhoneNumberVerified: req.IsPhoneNumberVerified,
		Gender:                req.Gender,
		Role:                  req.Role,
	}
	user, err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (s *service) GetUserByID(userID uint) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) GetUserByUsername(username string) (*models.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) UpdateUser(req *params.UpdateUserRequest) (*models.User, error) {
	user, err := s.userRepo.GetUserByID(req.ID)
	if err != nil {
		return nil, err
	}

	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.PhoneNumber = req.PhoneNumber
	user.Gender = req.Gender
	user.Avatar = req.Avatar

	user, err = s.userRepo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *service) DeleteUser(userID uint) error {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return err
	}
	if err := s.userRepo.DeleteUser(user); err != nil {
		return err
	}
	return nil
}
