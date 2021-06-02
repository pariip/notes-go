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

func (h *handler) createNote(c echo.Context) error {
	lang := getLanguage(c)

	user, err := h.getUserInJwtToken(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(lang, messages.ParseQueryError),
		}
	}

	req := new(params.CreateNoteRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.note",
			Function: "createNote",
			Message:  h.translator.TranslateEn(err.Error()),
		})

		message := messages.ParseQueryError

		if cerrors.As(err) {
			message = err.Error()
		}

		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(lang, message),
		}
	}
	req.UserID = user.ID

	note, err := h.noteService.CreateNote(req)
	if err != nil {
		message, code := cerrors.HttpError(err)
		h.logger.Error(&log.Field{
			Section:  "server.note",
			Function: "createNote",
			Params:   map[string]interface{}{"req": req},
			Message:  h.translator.TranslateEn(err.Error()),
		})
		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(lang, message),
		}

	}

	return c.JSON(http.StatusOK, note)
}

func (h *handler) getAllNotes(c echo.Context) error {
	lang := getLanguage(c)

	user, err := h.getUserInJwtToken(c)
	if err != nil {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(lang, messages.ParseQueryError),
		}
	}

	notes, err := h.noteService.GetAllNotes(user.ID)
	if err != nil {
		message, code := cerrors.HttpError(err)
		h.logger.Error(&log.Field{
			Section:  "server.note",
			Function: "getAllNotes",
			Params:   map[string]interface{}{"user_id": user.ID},
			Message:  h.translator.TranslateEn(err.Error()),
		})
		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(lang, message),
		}

	}

	return c.JSON(http.StatusOK, notes)
}

func (h *handler) getNoteByID(c echo.Context) error {
	lang := getLanguage(c)

	noteID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.note",
			Function: "getNoteByID",
			Params:   map[string]interface{}{"id": noteID},
			Message:  err.Error(),
		})
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(lang, messages.ParseQueryError),
		}
	}

	note, err := h.noteService.GetNoteByID(uint(noteID))
	if err != nil {
		message, code := cerrors.HttpError(err)
		return &echo.HTTPError{
			Code:     code,
			Message:  h.translator.Translate(lang, message),
			Internal: nil,
		}
	}

	return c.JSON(http.StatusOK, note)
}
