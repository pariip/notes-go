package picture

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
)

type service struct {
	cfg        *config.Path
	imageRepo  contract.ImageRepository
	logger     log.Logger
	translator translate.Translator
}

func New(cfg config.Path, mainRepo contract.MainRepository, logger log.Logger, translator translate.Translator) contract.ImageService {
	return &service{
		cfg:        &cfg,
		imageRepo:  mainRepo,
		logger:     logger,
		translator: translator,
	}
}
