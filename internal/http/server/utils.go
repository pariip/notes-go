package server

import (
	"github.com/labstack/echo/v4"
	"github.com/pariip/notes-go/pkg/translate"
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
