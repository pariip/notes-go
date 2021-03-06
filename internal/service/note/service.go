package note

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
)

type service struct {
	cfg        *config.Config
	noteRepo   contract.MainRepository
	validate   contract.NoteValidation
	logger     log.Logger
	translator translate.Translator
}

func New(cfg *config.Config, noteRepo contract.MainRepository, validate contract.NoteValidation, logger log.Logger, translator translate.Translator) contract.NoteService {
	return &service{
		cfg:        cfg,
		noteRepo:   noteRepo,
		validate:   validate,
		logger:     logger,
		translator: translator,
	}
}
