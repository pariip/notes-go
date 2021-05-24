package types

import (
	"fmt"
	"github.com/pariip/notes-go/pkg/cerrors"
	"github.com/pariip/notes-go/pkg/translate/messages"
)

type (
	NoteStatus uint
)

const (
	_ NoteStatus = iota
	NotePending
	NoteConfirmed
	NoteReject
	NoteClose
)

var (
	noteStatus = map[NoteStatus]string{
		NotePending:   "pending",
		NoteConfirmed: "confirmed",
		NoteReject:    "reject",
		NoteClose:     "close",
	}
)

func (n NoteStatus) String() string {
	if s, ok := noteStatus[n]; ok {
		return s
	}
	return fmt.Sprintf("NoteStatus(%d)", n)
}
func (n NoteStatus) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}
func (n *NoteStatus) UnmarshalText(by []byte) error {
	for i, v := range noteStatus {
		if v == string(by) {
			*n = i
			return nil
		}
	}
	return cerrors.New(cerrors.KindInvalid, messages.InvalidBookStatus)
}
