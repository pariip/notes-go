package schema

import (
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/models/types"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username              string `gorm:"unique;not null"`
	Password              string
	FirstName             string
	LastName              string
	Email                 string
	IsEmailVerified       bool
	PhoneNumber           string
	IsPhoneNumberVerified bool
	Gender                types.Gender
	Role                  types.Role
	Avatar                string
}

func (u *User) ConvertModel() *models.User {
	return &models.User{
		ID:                    u.ID,
		Username:              u.Username,
		Password:              u.Password,
		FirstName:             u.FirstName,
		LastName:              u.LastName,
		Email:                 u.Email,
		IsEmailVerified:       u.IsEmailVerified,
		PhoneNumber:           u.PhoneNumber,
		IsPhoneNumberVerified: u.IsPhoneNumberVerified,
		Gender:                u.Gender,
		Role:                  u.Role,
		Avatar:                u.Avatar,
	}
}

func ConvertUser(user *models.User) *User {
	return &User{
		Model:                 gorm.Model{ID: user.ID},
		Username:              user.Username,
		Password:              user.Password,
		FirstName:             user.FirstName,
		LastName:              user.LastName,
		Email:                 user.Email,
		IsEmailVerified:       user.IsEmailVerified,
		PhoneNumber:           user.PhoneNumber,
		IsPhoneNumberVerified: user.IsPhoneNumberVerified,
		Gender:                user.Gender,
		Role:                  user.Role,
		Avatar:                user.Avatar,
	}
}
