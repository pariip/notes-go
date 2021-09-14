package validation

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
)

type service struct {
	cfg        *config.Validation
	logger     log.Logger
	translator translate.Translator
}

func New(cfg *config.Validation, logger log.Logger, translator translate.Translator) contract.ValidationService {
	return &service{
		cfg:        cfg,
		logger:     logger,
		translator: translator,
	}
}
