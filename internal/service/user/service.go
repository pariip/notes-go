package user

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
)

type service struct {
	userCfg    *config.User
	userRepo   contract.UserRepository
	validate   contract.AuthValidation
	logger     log.Logger
	translator translate.Translator
}

func New(cfg config.User, mainRepo contract.MainRepository, validationService contract.AuthValidation, logger log.Logger, translator translate.Translator) contract.UserService {
	return &service{
		userCfg:    &cfg,
		userRepo:   mainRepo,
		validate:   validationService,
		logger:     logger,
		translator: translator,
	}
}
