package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pariip/notes-go/pkg/translate"
	"golang.org/x/text/language"
	"path/filepath"
)

type messageBundle struct {
	bundle *i18n.Bundle
}

func New(path string) (translate.Translator, error) {

	bundle := &messageBundle{
		bundle: i18n.NewBundle(language.English),
	}

	if err := bundle.loadBundle(path); err != nil {
		return nil, err
	}

	return bundle, nil
}

func (m *messageBundle) loadBundle(path string) error {
	m.bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	messageFiles, err := filepath.Glob(filepath.Join(path, "*.toml"))
	if err != nil {
		return err
	}
	for _, messageFile := range messageFiles {
		_, err := m.bundle.LoadMessageFile(messageFile)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *messageBundle) getLocalized(lang string) *i18n.Localizer {
	return i18n.NewLocalizer(m.bundle, lang)
}

func (m *messageBundle) Translate(lang translate.Language, key string) string {
	message, err := m.getLocalized(string(lang)).Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})
	if err != nil {
		return key
	}
	return message
}

func (m *messageBundle) TranslateEn(key string) string {
	message, err := m.getLocalized("en").Localize(&i18n.LocalizeConfig{
		MessageID: key,
	})
	if err != nil {
		return key
	}
	return message
}
