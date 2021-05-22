package app

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/pkg/log/logrus"
	"github.com/pariip/notes-go/pkg/translate/i18n"
)

func Run(cfg *config.Config) error {
	logger, err := logrus.New(&logrus.Option{
		Path:         cfg.Logger.InternalPath,
		Pattern:      cfg.Logger.FilenamePattern,
		MaxAge:       cfg.Logger.MaxAge,
		RotationTime: cfg.Logger.RotationTime,
		RotationSize: cfg.Logger.MaxSize,
	})

	if err != nil {
		return err
	}
	_ = logger
	translatorServ, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}
	_ = translatorServ
	return nil
}
