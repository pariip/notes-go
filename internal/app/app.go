package app

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/db/postgres"
	"github.com/pariip/notes-go/internal/http/server"
	"github.com/pariip/notes-go/internal/service/auth"
	"github.com/pariip/notes-go/internal/service/user"
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

	translatorServ, err := i18n.New(cfg.I18n.BundlePath)
	if err != nil {
		return err
	}

	mainRepository, err := postgres.New(cfg.Database.Postgres, translatorServ, logger)

	authService := auth.New(cfg.Auth, mainRepository, logger, translatorServ)
	userService := user.New(cfg.User, mainRepository, authService, logger, translatorServ)

	handler := server.NewHttpHandler(&server.HandlerFields{
		Cfg:         cfg,
		UserService: userService,
		Logger:      logger,
		Translator:  translatorServ,
	})
	httpServer := server.NewHttpServer(handler)

	return httpServer.Start(8003)

}
