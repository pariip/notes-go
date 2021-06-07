package server

import (
	"github.com/labstack/echo/v4"
	"github.com/pariip/notes-go/internal/params"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"net/http"
	"strconv"
)

func (h *handler) adminCreateUser(c echo.Context) error {
	lang := getLanguage(c)
	req := new(params.CreateUserRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "server",
			Function: "createUser",
			Message:  h.translator.Translate(err.Error()),
		})
		message := messages.ParseQueryError
		if cerrors.As(err) {
			message = err.Error()
		}
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: message,
		}
	}
	user, err := h.userService.CreateUser(req)
	if err != nil {
		message, code := cerrors.HttpError(err)
		h.logger.Error(&log.Field{
			Section:  "server",
			Function: "adminCreateUser",
			Params:   map[string]interface{}{"req": req},
			Message:  h.translator.Translate(err.Error()),
		})
		return &echo.HTTPError{
			Code:     code,
			Message:  h.translator.Translate(message, lang),
			Internal: nil,
		}
	}
	return c.JSON(http.StatusOK, user)
}

func (h *handler) adminGetUser(c echo.Context) error {
	lang := getLanguage(c)
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server",
			Function: "AdminGetUser",
			Params:   map[string]interface{}{"id": idString},
			Message:  "",
		})
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(messages.ParseQueryError, lang),
		}
	}
	user, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		message, code := cerrors.HttpError(err)

		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(message, lang),
		}
	}
	return c.JSON(http.StatusOK, user)

}

func (h *handler) adminUpdateUser(c echo.Context) error {
	lang := getLanguage(c)
	req := new(params.UpdateUserRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "server",
			Function: "adminUpdateUser",
			Message:  err.Error(),
		})
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(messages.ParseQueryError, lang),
		}
	}
	user, err := h.userService.UpdateUser(req)
	if err != nil {
		message, code := cerrors.HttpError(err)
		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(message, lang),
		}
	}
	return c.JSON(http.StatusOK, user)

}

func (h *handler) adminDeleteUser(c echo.Context) error {
	lang := getLanguage(c)

	userID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server",
			Function: "getUserByID",
			Params:   map[string]interface{}{"id": userID},
			Message:  err.Error(),
		})
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(messages.ParseQueryError, lang),
		}
	}

	err = h.userService.DeleteUser(uint(userID))
	if err != nil {
		message, code := cerrors.HttpError(err)

		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(message, lang),
		}
	}
	return c.NoContent(http.StatusOK)

}
