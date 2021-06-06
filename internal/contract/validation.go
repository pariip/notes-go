package contract

type (
	ValidationService interface {
		Username(username string) error
		Password(password string) error
	}
)
