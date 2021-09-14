package params

import "github.com/pariip/notes-go/internal/models/types"

type (
	SignupRequest struct {
		Username    string       `json:"username"`
		FirstName   string       `json:"first_name"`
		LastName    string       `json:"last_name"`
		Email       string       `json:"email"`
		PhoneNumber string       `json:"phone_number"`
		Gender      types.Gender `json:"gender"`
		Password    string       `json:"password"`
	}

	LoginRequest struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	UserTokens struct {
		AccessToken  string `yaml:"access_token"`
		RefreshToken string `yaml:"refresh_token"`
	}
)
