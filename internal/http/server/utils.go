package server

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type LangQ struct {
	lang string
	q    float64
}

func getLanguage(c echo.Context) []translate.Language {
	acceptLanguages := c.Request().Header.Get("Accept-Language")

	var lqs []LangQ
	languages := strings.Split(acceptLanguages, ",")

	for _, language := range languages {
		language = strings.Trim(language, " ")
		langWithQ := strings.Split(language, ";")

		if len(langWithQ) == 1 {
			lq := LangQ{lang: langWithQ[0], q: 1}
			lqs = append(lqs, lq)
		} else {
			valueQ := strings.Split(langWithQ[1], "=")
			q, err := strconv.ParseFloat(valueQ[1], 64)
			if err != nil {
				continue
			}
			lq := LangQ{langWithQ[0], q}
			lqs = append(lqs, lq)

		}
	}
	sort.SliceStable(lqs, func(i, j int) bool {
		return lqs[i].q > lqs[j].q
	})
	var result []translate.Language
	for _, lq := range lqs {
		result = append(result, translate.GetLanguage(lq.lang))
	}
	return result
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
			Message: h.translator.Translate(messages.InvalidToken, lang...),
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
			Message: h.translator.Translate(messages.InvalidToken, lang...),
		}
	}
	return user, nil
}
