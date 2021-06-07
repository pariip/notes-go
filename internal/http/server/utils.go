package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"net/http"
)

func getLanguage(c echo.Context) translate.Language {
	language := c.Request().Header.Get("Accept-Language")
	switch language {
	case "fa", "fa-ir", "farsi", "persian":
		return translate.FA
	case "en", "en-us", "english":
		return translate.EN
	default:
		return translate.EN
	}
}

func (h *handler) getUserInJwtToken(c echo.Context) (*models.Claims, error) {
	lang := getLanguage(c)
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		h.logger.Error(&log.Field{
			Section:  "server.note",
			Function: "getUserInJwtToken",
			Message:  h.translator.Translate(messages.InvalidToken),
		})

		return nil, &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: h.translator.Translate(messages.InvalidToken, lang),
		}
	}

	user, ok := token.Claims.(*models.Claims)
	if !ok {
		h.logger.Error(&log.Field{
			Section:  "server.note",
			Function: "getUserInJwtToken",
			Message:  h.translator.Translate(messages.InvalidToken),
		})

		return nil, &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: h.translator.Translate(messages.InvalidToken, lang),
		}
	}
	return user, nil
}
