package server

import (
	"github.com/labstack/echo/v4"
	"github.com/pariip/notes-go/internal/models/types"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"net/http"
)

func middlewarePermission(h *handler, roles ...types.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			lang := getLanguage(c)

			user, err := h.getUserInJwtToken(c)
			if err != nil {
				return &echo.HTTPError{
					Code:    http.StatusBadRequest,
					Message: h.translator.Translate(messages.ParseQueryError, lang...),
				}
			}

			for _, role := range roles {
				if user.Role == role {
					return next(c)
				}
			}

			h.logger.Error(&log.Field{
				Section:  "server.middleware",
				Function: "middlewarePermission",
				Message:  h.translator.Translate(messages.NotPermission),
			})

			return &echo.HTTPError{
				Code:    http.StatusForbidden,
				Message: h.translator.Translate(messages.NotPermission, lang...),
			}
		}
	}
}
