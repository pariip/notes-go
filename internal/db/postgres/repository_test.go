package postgres

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/db/postgres/schema"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/models/types"
	"github.com/pariip/notes-go/pkg/log/logrus"
	"github.com/pariip/notes-go/pkg/random"
	"github.com/pariip/notes-go/pkg/translate/i18n"
	"log"
	"testing"
)

var repositoryTest *repository

func setupTest(t *testing.T) {
	cfg := &config.Postgres{
		Username:  "postgres",
		Password:  "123456",
		DBName:    "go_notes_test",
		Host:      "127.0.0.1",
		Port:      "5432",
		SSLMode:   "disable",
		TimeZone:  "Asia/Tehran",
		Charset:   "utf8mb4",
		Migration: true,
	}

	translator, err := i18n.New("../../../build/i18n/")
	if err != nil {
		t.Fatal(err)
	}

	logger, err := logrus.New(&logrus.Option{
		Path:         "../../../logs/test",
		Pattern:      "%Y-%m-%dT%H:%M",
		MaxAge:       "720h",
		RotationTime: "24h",
		RotationSize: "20MB",
	})
	if err != nil {
		t.Fatal(err)
	}

	repositoryTest = &repository{
		cfg:        cfg,
		translator: translator,
		logger:     logger,
	}
	if err := repositoryTest.connect(); err != nil {
		t.Fatal(err)
	}
	if err := repositoryTest.db.Migrator().DropTable(
		&schema.User{},
		&schema.Note{},
		&schema.Picture{},
	); err != nil {
		log.Fatalln(err)
	}

	if err := repositoryTest.db.Migrator().CreateTable(
		&schema.User{},
		&schema.Note{},
		&schema.Picture{},
	); err != nil {
		log.Fatalln(err)
	}
}

func teardownTest() {
	repositoryTest = nil
}

func newUserTest() *models.User {
	return &models.User{
		Username:    random.String(8),
		Password:    random.String(25),
		FirstName:   random.String(8),
		LastName:    random.String(8),
		Email:       random.String(5) + "@" + random.String(3) + "." + random.String(3),
		PhoneNumber: "0918" + random.CreateStringWithCharset(7, "0123456789"),
		Gender:      types.Female,
		Role:        types.Admin,
		Avatar:      "",
	}
}

func newNoteTest(user *models.User) *models.Note {

	return &models.Note{
		UserID:      user.ID,
		Title:       random.String(8),
		Description: random.String(45),
		PublicNote:  false,
		Pictures:    nil,
	}

}
