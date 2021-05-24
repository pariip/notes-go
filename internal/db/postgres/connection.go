package postgres

import (
	"fmt"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/log"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func (r *repository) connect() error {
	dsn := fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=%v TimeZone=%v ",
		r.cfg.Username,
		r.cfg.Password,
		r.cfg.DBName,
		r.cfg.Host,
		r.cfg.Port,
		r.cfg.SSLMode,
		r.cfg.TimeZone,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		r.logger.Error(&log.Field{
			Section:  "postgres.connection",
			Function: "connection",
			Message:  err.Error(),
		})
		return cerrors.New(cerrors.KindUnexpected, messages.DBError)
	}
	r.db = db
	return nil
}
