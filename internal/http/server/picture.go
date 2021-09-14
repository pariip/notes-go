package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"net/http"
)

func (h *handler) uploadImage(c echo.Context) error {
	lang := getLanguage(c)

	image, err := c.FormFile("image")
	if err != nil {
		h.logger.Error(&log.Field{
			Section:  "server.picture",
			Function: "imageUpload",
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

	upload, err := h.imageService.UploadImage(image)
	if err != nil {
		message, code := cerrors.HttpError(err)
		h.logger.Error(&log.Field{
			Section:  "server.picture",
			Function: "imageUpload",
			Params:   map[string]interface{}{"imageName": image.Filename},
			Message:  h.translator.Translate(err.Error()),
		})
		return &echo.HTTPError{
			Code:    code,
			Message: h.translator.Translate(message, lang...),
		}
	}

	fmt.Printf("%T")
	fmt.Println(image.Filename, image.Size)
	return c.JSON(http.StatusOK, upload)
}
