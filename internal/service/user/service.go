package user

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
)

type service struct {
	userCfg     *config.User
	userRepo    contract.UserRepository
	authService contract.AuthService
	logger      log.Logger
	translator  translate.Translator
}

func New(cfg config.User, mainRepo contract.MainRepository, authService contract.AuthService, logger log.Logger, translator translate.Translator) contract.UserService {
	return &service{
		userCfg:     &cfg,
		userRepo:    mainRepo,
		authService: authService,
		logger:      logger,
		translator:  translator,
	}
}
