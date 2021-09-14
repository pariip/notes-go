package auth

import (
	"github.com/golang/mock/gomock"
	"github.com/pariip/notes-go/internal/config"
	"github.com/pariip/notes-go/internal/mock/repository_mock"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/models/types"
	"github.com/pariip/notes-go/pkg/log/logrus"
	"github.com/pariip/notes-go/pkg/random"
	"github.com/pariip/notes-go/pkg/translate/i18n"
	"math/rand"
	"testing"
)

var (
	mockCtrl     *gomock.Controller
	mockMainRepo *repository_mock.MockUserRepository
	serviceTest  *service
)

func setupTest(t *testing.T) {
	cfg := &config.Auth{
		AccessExpirationInMinute:  15,
		RefreshExpirationInMinute: 525600,
		JWTSecret:                 "secret",
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

	mockCtrl = gomock.NewController(t)
	mockMainRepo = repository_mock.NewMockUserRepository(mockCtrl)

	serviceTest = &service{
		cfg:        cfg,
		userRepo:   mockMainRepo,
		logger:     logger,
		translator: translator,
	}
}

func teardownTest() {
	mockCtrl.Finish()
	mockCtrl = nil
	mockMainRepo = nil
	serviceTest = nil
}

func newUserTest() *models.User {
	return &models.User{
		ID:          uint(rand.Uint32()),
		Username:    random.String(8),
		Password:    random.String(25),
		FirstName:   random.String(8),
		LastName:    random.String(8),
		Email:       random.String(5) + "@" + random.String(3) + "." + random.String(3),
		PhoneNumber: "0912" + random.CreateStringWithCharset(7, "0123456789"),
		Gender:      types.Male,
		Role:        types.Basic,
		Avatar:      "",
	}
}
