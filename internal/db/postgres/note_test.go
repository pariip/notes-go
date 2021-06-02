package postgres

import (
	"errors"
	"fmt"
	"github.com/pariip/notes-go/internal/models"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/random"
	"github.com/pariip/notes-go/pkg/translate/messages"
	"math/rand"
	"testing"
)

func newNoteTest(user *models.User) *models.Note {

	return &models.Note{
		UserID:      user.ID,
		Title:       random.String(8),
		Description: random.String(45),
		PublicNote:  false,
		Pictures:    nil,
	}

}

func Test_repository_CreateNote(t *testing.T) {
	repo := setupTest(t)
	user := newUserTest()
	var err error
	t.Run("create new user", func(t *testing.T) {
		if user, err = repo.CreateUser(user); err != nil {
			t.Fail()
		}
	})
	note := newNoteTest(user)
	tests := []struct {
		name string
		note *models.Note
		want error
	}{
		// TODO: Add test cases.
		{
			name: "create new note",
			note: note,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err = repo.CreateNote(tt.note)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}
}

func TestRepository_GetNoteByID(t *testing.T) {
	repo := setupTest(t)
	user := newUserTest()
	var err error
	t.Run("create new user", func(t *testing.T) {
		if user, err = repo.CreateUser(user); err != nil {
			t.Fail()
		}
	})
	note := newNoteTest(user)
	t.Run("create new note", func(t *testing.T) {
		if note, err = repo.CreateNote(note); err != nil {
			t.Fail()
		}
	})
	fmt.Println(note)
	tests := []struct {
		name string
		id   uint
		want error
	}{
		{
			name: "get note by id",
			id:   note.ID,
			want: nil,
		},
		{
			name: "note not found",
			id:   uint(rand.Uint64()),
			want: cerrors.New(cerrors.KindNotFound, messages.NoteNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err = repo.GetNoteByID(tt.id)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}

}

func TestRepository_GetAllNotes(t *testing.T) {
	repo := setupTest(t)
	user := newUserTest()
	var err error
	t.Run("create new user", func(t *testing.T) {
		if user, err = repo.CreateUser(user); err != nil {
			t.Fail()
		}
	})
	note := newNoteTest(user)
	t.Run("create new note", func(t *testing.T) {
		if note, err = repo.CreateNote(note); err != nil {
			t.Fail()
		}
	})
	tests := []struct {
		name   string
		userID uint
		want   error
	}{
		{
			name:   "get all note for userID",
			userID: note.UserID,
			want:   nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err = repo.GetAllNotes(tt.userID)
			if !errors.Is(err, tt.want) {
				t.Error()
			}
		})
	}

}

func TestRepository_UpdateNote(t *testing.T) {
	repo := setupTest(t)
	user := newUserTest()
	var err error

	t.Run("create new user", func(t *testing.T) {
		if user, err = repo.CreateUser(user); err != nil {
			t.Fail()
		}
	})

	note := newNoteTest(user)
	note1 := newNoteTest(user)
	t.Run("create new note", func(t *testing.T) {
		if note, err = repo.CreateNote(note); err != nil {
			t.Fail()
		}
	})

	tests := []struct {
		name string
		note *models.Note
		want error
	}{
		// TODO: Add test cases.
		{
			name: "update note",
			note: note,
			want: nil,
		},
		{
			name: "note not found",
			note: note1,
			want: cerrors.New(cerrors.KindNotFound, messages.NoteNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.UpdateNote(tt.note)
			if !errors.Is(err, tt.want) {
				t.Error()
			}

		})
	}
}

func TestRepository_DeleteNote(t *testing.T) {
	repo := setupTest(t)
	user := newUserTest()
	var err error

	t.Run("create new user", func(t *testing.T) {
		if user, err = repo.CreateUser(user); err != nil {
			t.Fail()
		}
	})

	note := newNoteTest(user)
	note1 := newNoteTest(user)
	t.Run("create new note", func(t *testing.T) {
		if note, err = repo.CreateNote(note); err != nil {
			t.Fail()
		}
	})

	tests := []struct {
		name string
		note *models.Note
		want error
	}{
		// TODO: Add test cases.
		{
			name: "update note",
			note: note,
			want: nil,
		},
		{
			name: "note not found",
			note: note1,
			want: cerrors.New(cerrors.KindNotFound, messages.NoteNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.DeleteNote(tt.note)
			if !errors.Is(err, tt.want) {
				t.Error()
			}

		})
	}
}
