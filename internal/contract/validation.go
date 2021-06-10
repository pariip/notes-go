package contract

//go:generate mockgen -source ./validation.go -package validation_mock -destination ../mock/validation_mock/mock.go

type (
	ValidationService interface {
		Username(username string) error
		Password(password string) error
	}
)
