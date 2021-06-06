package auth

import (
	"github.com/pariip/notes-go/internal/params"
	"github.com/pariip/notes-go/pkg/hash"
	"github.com/pariip/notes-go/pkg/log"
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
			Section:  "user.user_auth",
			Function: "Login",
			Params:   map[string]interface{}{"username": req.Username},
			Message:  s.translator.TranslateEn(err.Error()),
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
