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
			Message: h.translator.Translate(messages.ParseQueryError, lang...),
		}
	}

	req := new(params.CreateNoteRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.note",
			Function: "createNote",
			Message:  h.translator.Translate(err.Error()),
		})

		message := messages.ParseQueryError

		if cerrors.As(err) {
			message = err.Error()
		}

		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(message, lang...),
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
			Message:  h.translator.Translate(err.Error()),
		})
		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(message, lang...),
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
			Message: h.translator.Translate(messages.ParseQueryError, lang...),
		}
	}

	notes, err := h.noteService.GetAllNotes(user.ID)
	if err != nil {
		message, code := cerrors.HttpError(err)
		h.logger.Error(&log.Field{
			Section:  "server.note",
			Function: "getAllNotes",
			Params:   map[string]interface{}{"user_id": user.ID},
			Message:  h.translator.Translate(err.Error()),
		})
		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(message, lang...),
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
			Message: h.translator.Translate(messages.ParseQueryError, lang...),
		}
	}

	note, err := h.noteService.GetNoteByID(uint(noteID))
	if err != nil {
		message, code := cerrors.HttpError(err)
		return &echo.HTTPError{
			Code:     code,
			Message:  h.translator.Translate(message, lang...),
			Internal: nil,
		}
	}

	return c.JSON(http.StatusOK, note)
}

func (h *handler) updateNote(c echo.Context) error {
	lang := getLanguage(c)

	req := new(params.UpdateNoteRequest)
	if err := c.Bind(req); err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.note",
			Function: "updateNote",
			Message:  err.Error(),
		})
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(messages.ParseQueryError, lang...),
		}
	}

	note, err := h.noteService.UpdateNote(req)
	if err != nil {
		message, code := cerrors.HttpError(err)
		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(message, lang...),
		}
	}

	return c.JSON(http.StatusOK, note)
}

func (h *handler) deleteNote(c echo.Context) error {
	lang := getLanguage(c)

	noteID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.note",
			Function: "getUserByID",
			Params:   map[string]interface{}{"id": noteID},
			Message:  err.Error(),
		})
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: h.translator.Translate(messages.ParseQueryError, lang...),
		}
	}

	err = h.noteService.DeleteNote(uint(noteID))
	if err != nil {
		message, code := cerrors.HttpError(err)

		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(message, lang...),
		}
	}

	return c.NoContent(http.StatusOK)
}
