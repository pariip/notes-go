package user

import (
	"github.com/pariip/notes-go/internal/params"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/hash"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

func (s *service) Login(req *params.LoginRequest) (*params.UserTokens, error) {

	if err := s.validateUsername(req.Username); err != nil {
		return nil, err
	}

	if err := s.validatePassword(req.Password); err != nil {
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

	accessToken, err := s.authService.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	refreshToken, err := s.authService.GenerateRefreshToken(user)
	if err != nil {
		return nil, err
	}

	return &params.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *service) RefreshToken(refreshToken string, userID uint) (*params.UserTokens, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	ok, err := s.authService.RefreshTokenIsValid(refreshToken, userID)

	if err != nil {
		return nil, err
	}

	if !ok {
		s.logger.Error(&log.Field{
			Section:  "service.user",
			Function: "RefreshToken",
			Params:   map[string]interface{}{"user_id": userID},
			Message:  s.translator.TranslateEn(messages.InvalidToken),
		})
		return nil, cerrors.New(cerrors.KindInvalid, messages.InvalidToken)
	}

	accessToken, err := s.authService.GenerateAccessToken(user)
	if err != nil {
		return nil, err
	}

	return &params.UserTokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}
