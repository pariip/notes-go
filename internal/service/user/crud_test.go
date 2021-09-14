package user

import (
	"errors"
	"fmt"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/internal/params"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/hash"
	"github.com/pariip/notes-go/pkg/random"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"math/rand"
	"testing"
)

func TestCreateUser(t *testing.T) {
	setupTest(t)
	defer teardownTest()
	req := newCreateUserRequestTest()
	//req2 := newCreateUserRequestTest()

	//req.Password = "Sdwlp$2a15&"

	var password string
	var err error
	t.Run("hash password", func(t *testing.T) {
		password, err = hash.Password(req.Password)
		if err != nil {
			t.Fatal()
		}
	})

	user := &models.User{
		ID:                    1,
		Username:              req.Username,
		Password:              password,
		FirstName:             req.FirstName,
		LastName:              req.LastName,
		Email:                 req.Email,
		IsEmailVerified:       req.IsEmailVerified,
		PhoneNumber:           req.PhoneNumber,
		IsPhoneNumberVerified: req.IsPhoneNumberVerified,
		Gender:                req.Gender,
		Role:                  req.Role,
	}

	type validateUsername struct {
		expect    bool
		parameter string
		return1   error
	}
	type validatePass struct {
		expect    bool
		parameter string
		return1   error
	}
	type getUsername struct {
		expect    bool
		parameter string
		return1   bool
		return2   error
	}

	type createUser struct {
		expect    bool
		parameter *models.User
		return1   *models.User
		return2   error
	}

	test := []struct {
		name             string
		req              *params.CreateUserRequest
		validateUsername validateUsername
		validatePass     validatePass
		getUsername      getUsername
		createUser       createUser
		want             error
	}{
		{
			name: "correct test",
			req:  req,
			getUsername: getUsername{
				expect:    true,
				parameter: user.Username,
				return1:   true,
				return2:   nil,
			},
			validateUsername: validateUsername{
				expect:    true,
				parameter: user.Username,
				return1:   nil,
			},
			validatePass: validatePass{
				expect:    true,
				parameter: req.Password,
				return1:   nil,
			},
			createUser: createUser{
				expect:    true,
				parameter: user,
				return1:   user,
				return2:   nil,
			},
			want: nil,
		},
		//{
		//	name: "error happen in create",
		//	req:  req,
		//	createUser: createUser{
		//		expect:    true,
		//		parameter: user,
		//		return1:   nil,
		//		return2:   cerrors.New(cerrors.KindUnexpected, messages.DBError),
		//	},
		//	want: cerrors.New(cerrors.KindUnexpected, messages.DBError),
		//},
		//{
		//	name: "error happen in validate password",
		//	req:  req2,
		//	createUser: createUser{
		//		expect: false,
		//	},
		//	want: cerrors.New(cerrors.KindInvalid, messages.InvalidPassword),
		//},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.getUsername.expect {
				mockMainRepo.EXPECT().IsUsernameExist(tt.getUsername.parameter).Return(tt.getUsername.return1, tt.getUsername.return2)
			}
			if tt.createUser.expect {
				mockMainRepo.EXPECT().CreateUser(tt.createUser.parameter).Return(tt.createUser.return1, tt.createUser.return2)
			}
			if tt.validateUsername.expect {
				mockValidationService.EXPECT().Username(tt.validateUsername.parameter).Return(tt.validateUsername.return1)
			}
			if tt.validatePass.expect {
				mockValidationService.EXPECT().Password(tt.validatePass.parameter).Return(tt.validatePass.return1)
			}
			_, err := serviceTest.CreateUser(tt.req)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}

func TestGetUserByID(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()
	userID := uint(rand.Uint32())
	fmt.Println(user, userID)
	type getUser struct {
		expect    bool
		parameter uint
		return1   *models.User
		return2   error
	}
	tests := []struct {
		name    string
		userID  uint
		getUser getUser
		want    error
	}{
		// TODO: Add test cases.
		{
			name:   "correct test",
			userID: userID,
			getUser: getUser{
				expect:    true,
				parameter: userID,
				return1:   user,
				return2:   nil,
			},
			want: nil,
		},
		{
			name:   "invalid userID",
			userID: userID,
			getUser: getUser{
				expect:    true,
				parameter: userID,
				return1:   nil,
				return2:   cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
			},
			want: cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.getUser.expect {
				mockMainRepo.EXPECT().GetUserByID(tt.getUser.parameter).Return(tt.getUser.return1, tt.getUser.return2)
			}
			_, err := serviceTest.GetUserByID(tt.userID)
			if !errors.Is(err, tt.want) {
				t.Fail()
			}
		})
	}
}

func TestGetUserByUsername(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()
	username := random.String(5)

	type getUser struct {
		expect    bool
		parameter string
		return1   *models.User
		return2   error
	}

	test := []struct {
		name     string
		username string
		getUser  getUser
		want     error
	}{
		{
			name:     "correct test",
			username: username,
			getUser: getUser{
				expect:    true,
				parameter: username,
				return1:   user,
				return2:   nil,
			},
			want: nil,
		},
		{
			name:     "invalid username",
			username: username,
			getUser: getUser{
				expect:    true,
				parameter: username,
				return1:   nil,
				return2:   cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
			},
			want: cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.getUser.expect {
				mockMainRepo.EXPECT().GetUserByUsername(tt.getUser.parameter).Return(tt.getUser.return1, tt.getUser.return2)
			}
			_, err := serviceTest.GetUserByUsername(tt.username)
			if !errors.Is(err, tt.want) {
				t.Fail()
			}
		})
	}
}

func TestUpdateUser(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()
	req := newUpdateUserRequestTest()

	type getUser struct {
		expect    bool
		parameter uint
		return1   *models.User
		return2   error
	}

	type updateUser struct {
		expect    bool
		parameter *models.User
		return1   *models.User
		return2   error
	}

	test := []struct {
		name       string
		req        *params.UpdateUserRequest
		getUser    getUser
		updateUser updateUser
		want       error
	}{
		{
			name: "no error",
			req:  req,
			getUser: getUser{
				expect:    true,
				parameter: req.ID,
				return1:   user,
				return2:   nil,
			},
			updateUser: updateUser{
				expect:    true,
				parameter: user,
				return1:   user,
				return2:   nil,
			},
			want: nil,
		},
		{
			name: "error happen in get",
			req:  req,
			getUser: getUser{
				expect:    true,
				parameter: req.ID,
				return1:   nil,
				return2:   cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
			},
			updateUser: updateUser{
				expect: false,
			},
			want: cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
		},
		{
			name: "error happen in update",
			req:  req,
			getUser: getUser{
				expect:    true,
				parameter: req.ID,
				return1:   user,
				return2:   nil,
			},
			updateUser: updateUser{
				expect:    true,
				parameter: user,
				return1:   nil,
				return2:   cerrors.New(cerrors.KindUnexpected, messages.DBError),
			},
			want: cerrors.New(cerrors.KindUnexpected, messages.DBError),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.getUser.expect {
				mockMainRepo.EXPECT().GetUserByID(tt.getUser.parameter).Return(tt.getUser.return1, tt.getUser.return2)
			}
			if tt.updateUser.expect {
				mockMainRepo.EXPECT().UpdateUser(tt.updateUser.parameter).Return(tt.updateUser.return1, tt.updateUser.return2)
			}
			_, err := serviceTest.UpdateUser(tt.req)
			if !errors.Is(err, tt.want) {
				t.Fail()
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()
	userID := uint(rand.Uint32())

	type getUser struct {
		expect    bool
		parameter uint
		return1   *models.User
		return2   error
	}

	type deleteUser struct {
		expect    bool
		parameter *models.User
		return1   error
	}

	test := []struct {
		name       string
		userID     uint
		getUser    getUser
		deleteUser deleteUser
		want       error
	}{
		{
			name:   "no error",
			userID: userID,
			getUser: getUser{
				expect:    true,
				parameter: userID,
				return1:   user,
				return2:   nil,
			},
			deleteUser: deleteUser{
				expect:    true,
				parameter: user,
				return1:   nil,
			},
			want: nil,
		},
		{
			name:   "error happen in get",
			userID: userID,
			getUser: getUser{
				expect:    true,
				parameter: userID,
				return1:   nil,
				return2:   cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
			},
			deleteUser: deleteUser{
				expect: false,
			},
			want: cerrors.New(cerrors.KindNotFound, messages.UserNotFound),
		},
		{
			name:   "error happen in delete",
			userID: userID,
			getUser: getUser{
				expect:    true,
				parameter: userID,
				return1:   user,
				return2:   nil,
			},
			deleteUser: deleteUser{
				expect:    true,
				parameter: user,
				return1:   cerrors.New(cerrors.KindUnexpected, messages.DBError),
			},
			want: cerrors.New(cerrors.KindUnexpected, messages.DBError),
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			if tt.getUser.expect {
				mockMainRepo.EXPECT().GetUserByID(tt.getUser.parameter).Return(tt.getUser.return1, tt.getUser.return2)
			}
			if tt.deleteUser.expect {
				mockMainRepo.EXPECT().DeleteUser(tt.deleteUser.parameter).Return(tt.deleteUser.return1)
			}
			err := serviceTest.DeleteUser(tt.userID)
			if !errors.Is(err, tt.want) {
				t.Fail()
			}
		})
	}
}
