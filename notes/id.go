package notes

import "time"

const (
	// See https://pkg.go.dev/time#Layout
	noteIdTimeLayout = "20060102150405"
)

type noteId struct {
	time  time.Time
	value string
}

func (id noteId) Time() time.Time {
	return id.time
}

func (id noteId) Value() string {
	return id.value
}

func NewNoteId(t time.Time) noteId {
	return noteId{
		time:  t,
		value: t.Format(noteIdTimeLayout),
	}
}
