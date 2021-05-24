package postgres

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/contract"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate"
	"gorm.io/gorm"
)

type repository struct {
	db         *gorm.DB
	cfg        *config.Postgres
	translator translate.Translator
	logger     log.Logger
}

func New(cfg config.Postgres, translator translate.Translator, logger log.Logger) (contract.MainRepository, error) {

	repo := &repository{
		cfg:        &cfg,
		translator: translator,
		logger:     logger,
	}
	if err := repo.connect(); err != nil {
		return nil, err
	}
	if err := repo.migration(); err != nil {
		return nil, err
	}
	return repo, nil

}
