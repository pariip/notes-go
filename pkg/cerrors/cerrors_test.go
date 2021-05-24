package cerrors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestHttpError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		// TODO: Add test cases.
		{
			name:  "server error",
			args:  args{err: New(KindNotFound, "this page note found")},
			want:  "this page note found",
			want1: http.StatusNotFound,
		},
		{
			name: "server error",
			args: args{
				err: fmt.Errorf("something: %w", New(KindNotFound, "page not found")),
			},
			want:  "page not found",
			want1: http.StatusNotFound,
		},
		{
			name: "other errors",
			args: args{
				err: errors.New("something"),
			},
			want:  "something",
			want1: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := HttpError(tt.args.err)
			if got != tt.want {
				t.Errorf("HttpError() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("HttpError() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
