package auth

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
)

type service struct {
	cfg        *config.Auth
	authRepo   contract.AuthRepository
	logger     log.Logger
	translator translate.Translator
}

func New(cfg config.Auth, authRepo contract.AuthRepository, logger log.Logger, translator translate.Translator) contract.AuthService {
	return &service{
		cfg:        &cfg,
		authRepo:   authRepo,
		logger:     logger,
		translator: translator,
	}
}
