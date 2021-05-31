package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"time"
)

func (s service) GenerateAccessToken(user *models.User) (string, error) {
	accessExpirationTime := time.Now().Add(time.Duration(s.cfg.AccessExpirationInMinute) * time.Minute)

	claims := &models.Claims{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: accessExpirationTime.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := jwtToken.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		s.logger.Error(&log.Field{
			Section:  "auth.auth",
			Function: "GenerateAccessToken",
			Params:   map[string]interface{}{"user": user},
			Message:  err.Error(),
		})
		return "", cerrors.New(cerrors.KindUnexpected, messages.GeneralError)

	}
	return accessToken, nil
}

func (s service) GenerateRefreshToken(user *models.User) (string, error) {
	refreshExpirationTime := time.Now().Add(time.Duration(s.cfg.RefreshExpirationInMinute) * time.Minute)

	claims := &models.Claims{
		ID:          user.ID,
		Username:    user.Username,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Role:        user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: refreshExpirationTime.Unix(),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := jwtToken.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		s.logger.Error(&log.Field{
			Section:  "auth.auth",
			Function: "GenerateRefreshToken",
			Params:   map[string]interface{}{"user": user},
			Message:  err.Error(),
		})
		return "", cerrors.New(cerrors.KindUnexpected, messages.GeneralError)
	}
	if err := s.authRepo.CreateToken(refreshToken, user.ID); err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (s *service) RefreshTokenIsValid(token string, userID uint) (bool, error) {
	return s.authRepo.TokenIsExistWithUserID(token, userID)
}
