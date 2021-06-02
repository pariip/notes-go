package note

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
)

type service struct {
	cfg        *config.Config
	noteRepo   contract.NoteRepository
	logger     log.Logger
	translator translate.Translator
}

func New(cfg *config.Config, noteRepo contract.MainRepository, logger log.Logger, translator translate.Translator) contract.NoteService {
	return &service{
		cfg:        cfg,
		noteRepo:   noteRepo,
		logger:     logger,
		translator: translator,
	}
}
