package i18n

import (
	"github.com/pariip/notes-go/pkg/translate"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"testing"
)

func Test_messageBundle_Translate(t *testing.T) {

	type args struct {
		language translate.Language
		message  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "translate farsi",
			args: args{
				message:  messages.UserNotFound,
				language: translate.GetLanguage("fa"),
			},
			want: "کاربر مورد نظر یافت نشد",
		},
		{
			name: "translate english",
			args: args{
				message:  messages.UserNotFound,
				language: translate.GetLanguage("en"),
			},
			want: "UserNotFound",
		},
		{
			name: "message key not found",
			args: args{
				message:  "NoKeyFound",
				language: translate.GetLanguage("en"),
			},
			want: "NoKeyFound",
		},
	}
	translate, err := New("../../../build/i18n/")
	if err != nil {
		t.Errorf("New() error : %v", err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translate.Translate(tt.args.language, tt.args.message)
			if got != tt.want {
				t.Errorf("Translate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMessageBundle_TranslateEn(t *testing.T) {

	type args struct {
		message string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "translate",
			args: args{
				message: messages.UserNotFound,
			},
			want: "UserNotFound",
		},
		{
			name: "message key not found",
			args: args{
				message: "NoKeyFound",
			},
			want: "NoKeyFound",
		},
	}

	translate, err := New("../../../build/i18n/")
	if err != nil {
		t.Errorf("New() error : %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translate.TranslateEn(tt.args.message)
			if got != tt.want {
				t.Errorf("TranslateEn() got = %v, want %v", got, tt.want)
			}

		})
	}
}
