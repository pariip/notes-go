package translate

import "strings"

type (
	Translator interface {
		Translate(key string, lang ...Language) string
	}
	Language string
)

const (
	EN = "en"
	FA = "fa"
)

func GetLanguage(lang string) Language {
	switch strings.ToLower(lang) {
	case "en":
		return EN
	case "fa":
		return FA
	default:
		return EN
	}
}
