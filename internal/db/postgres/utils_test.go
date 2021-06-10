package postgres

import (
	"math/rand"
	"testing"
)

func TestIsErrorNotFound(t *testing.T) {
	setupTest(t)
	defer teardownTest()

	user := newUserTest()
	var err error

	t.Run("create new record", func(t *testing.T) {
		if user, err = repositoryTest.CreateUser(user); err != nil {
			t.Fail()
		}
	})

	err1 := repositoryTest.db.First(user, rand.Uint32()).Error
	err2 := repositoryTest.db.First(user, user.ID).Error

	test := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "error not found",
			err:  err1,
			want: true,
		},
		{
			name: "error nil",
			err:  err2,
			want: false,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			err := isErrorNotFound(tt.err)
			if err != tt.want {
				t.Error()
			}
		})
	}
}
