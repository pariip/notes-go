package hash

import (
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/random"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"testing"
)

func TestCheckPassword(t *testing.T) {
	password := random.String(6)
	hashPassword, err := Password(password)
	if err != nil {
		t.Fail()
	}
	tests := []struct {
		name string
		pass string
		hash string
		want error
	}{
		{
			name: "correct test case",
			pass: password,
			hash: hashPassword,
			want: nil,
		}, {
			name: "incorrect test case",
			pass: random.String(7),
			hash: hashPassword,
			want: cerrors.New(cerrors.KindInvalid, messages.UsernameOrPasswordIsIncorrect),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckPassword(tt.pass, tt.hash)
			if err != tt.want {
				t.Fail()
			}
		})
	}
}
