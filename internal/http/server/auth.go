package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/pariip/notes-go/internal/params"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"net/http"
	"strconv"
)

func (h *handler) login(c echo.Context) error {
	lang := getLanguage(c)

	req := new(params.LoginRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "http.server",
			Function: "login",
			Message:  err.Error(),
		})

		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(messages.ParseQueryError, lang...),
		}
	}

	tokens, err := h.authService.Login(req)
	if err != nil {
		message, code := cerrors.HttpError(err)
		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(message, lang...),
		}
	}

	return c.JSON(http.StatusOK, tokens)
}

func (h *handler) signup(c echo.Context) error {
	lang := getLanguage(c)

	req := new(params.SignupRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "http.server",
			Function: "signup",
			Params:   map[string]interface{}{"req": req},
			Message:  err.Error(),
		})
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(messages.ParseQueryError, lang...),
		}
	}

	token, err := h.authService.Signup(req)
	if err != nil {
		message, code := cerrors.HttpError(err)

		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(message, lang...),
		}

	}
	return c.JSON(http.StatusOK, token)
}

func (h *handler) refreshToken(c echo.Context) error {
	lang := getLanguage(c)

	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		h.logger.Error(&log.Field{
			Section:  "server.auth",
			Function: "refreshToken",
			Message:  h.translator.Translate(messages.InvalidToken),
		})

		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: h.translator.Translate(messages.InvalidToken, lang...),
		}
	}

	userIdString := c.Param("id")
	userID, err := strconv.Atoi(userIdString)
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.auth",
			Function: "refreshToken",
			Params:   map[string]interface{}{"user_id_string": userIdString},
			Message:  err.Error(),
		})

		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(messages.ParseQueryError, lang...),
		}
	}

	res, err := h.authService.RefreshToken(token.Raw, uint(userID))
	if err != nil {
		message, code := cerrors.HttpError(err)

		return &echo.HTTPError{
			Code:    code,
			Message: message,
		}
	}

	return c.JSON(http.StatusOK, res)
}
