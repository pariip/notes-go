package auth

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/models/types"
	"github.com/pariip/notes-go/internal/params"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/hash"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

func (s *service) Login(req *params.LoginRequest) (*params.UserTokens, error) {

	if err := s.validate.Username(req.Username); err != nil {
		return nil, err
	}

	if err := s.validate.Password(req.Password); err != nil {
		return nil, err
	}

	user, err := s.userRepo.GetUserByUsername(req.Username)
	if err != nil {
		return nil, err
	}

	if err := hash.CheckPassword(req.Password, user.Password); err != nil {
		s.logger.Error(&log.Field{
			Section:  "server.auth",
			Function: "Login",
			Params:   map[string]interface{}{"username": req.Username},
			Message:  s.translator.Translate(err.Error()),
		})

		return nil, err
	}

	accessToken, err := s.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return &params.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *service) Signup(req *params.SignupRequest) (*params.UserTokens, error) {
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
			Section:  "service.auth",
			Function: "Register",
			Params:   map[string]interface{}{"username": req.Username},
			Message:  s.translator.Translate(messages.UsernameIsDuplicate),
		})
		return nil, cerrors.New(cerrors.KindInvalid, messages.UsernameIsDuplicate)
	}

	password, err := hash.Password(req.Password)
	if err != nil {
		s.logger.Error(&log.Field{
			Section:  "service.auth",
			Function: "Signup",
			Params:   map[string]interface{}{"username": req.Username},
			Message:  s.translator.Translate(err.Error()),
		})
		return nil, err
	}
	user := &models.User{
		Username:    req.Username,
		Password:    password,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Gender:      req.Gender,
		Role:        types.Basic,
		Avatar:      "",
	}

	user, err = s.userRepo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}
	return &params.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
