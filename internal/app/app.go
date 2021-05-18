package app

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/pkg/translate/i18n"
)

func Run(cfg *config.Config) error {
	translatorServ, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}
	_ = translatorServ
	return nil
}
