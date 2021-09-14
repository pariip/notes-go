package auth

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
)

type service struct {
	cfg        *config.Auth
	userRepo   contract.UserRepository
	validate   contract.ValidationService
	logger     log.Logger
	translator translate.Translator
}

func New(cfg config.Auth, userRepo contract.UserRepository, validationService contract.ValidationService, logger log.Logger, translator translate.Translator) contract.AuthService {
	return &service{
		cfg:        &cfg,
		userRepo:   userRepo,
		validate:   validationService,
		logger:     logger,
		translator: translator,
	}
}
