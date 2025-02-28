package notes

import "time"

const (
	// See https://pkg.go.dev/time#Layout
	noteIdTimeLayout = "20060102150405"
)

type NoteId struct {
	time  time.Time
	value string
}

func (id NoteId) Time() time.Time {
	return id.time
}

func (id NoteId) Value() string {
	return id.value
}

func NewNoteId(t time.Time) *NoteId {
	return &NoteId{
		time:  t,
		value: t.Format(noteIdTimeLayout),
	}
}
