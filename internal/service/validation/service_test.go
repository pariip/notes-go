package validation

import (
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/pkg/log/logrus"
	"github.com/pariip/notes-go/pkg/translate/i18n"
	"testing"
)

var serviceTest *service

func setupTest(t *testing.T) {
	cfg := &config.Validation{
		UsernameMinLength:  3,
		UsernameMaxLength:  50,
		PasswordMinLetters: 4,
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
	serviceTest = &service{
		cfg:        cfg,
		logger:     logger,
		translator: translator,
	}
}
func teardownTest() {
	serviceTest = nil
}
