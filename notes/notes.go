package notes

import "errors"

type NoteType string

const (
	TextNoteType       NoteType = "text"
	ImageNoteType      NoteType = "image"
	LinkNoteType       NoteType = "link"
	VideoNoteType      NoteType = "video"
	AudioNoteType      NoteType = "audio"
	PdfNoteType        NoteType = "pdf"
	ExcalidrawNoteType NoteType = "excalidraw"
	FileNoteType       NoteType = "file"
	CodeNoteType       NoteType = "code"
	QuoteNoteType      NoteType = "quote"
	LatexNoteType      NoteType = "latex"
)

var (
	ErrInvalidNoteType = errors.New("invalid note type")
)

type NoteMetadata struct {
	Id       string
	Type     NoteType
	Content  string
	Path     []string
	Children []string
}

func ReadCurrentNote() {}

func ChangeCurrentNote() {}

func NestNote() {}

func UnnestNote() {}
