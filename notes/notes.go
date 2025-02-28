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

type Note struct {
	Id       string
	Type     NoteType
	Path     []string
	Children []string
}

type TextNote struct {
	Note    Note
	Content string
}

type ImageNote struct {
	Note    Note
	Content string
}

type LinkNote struct {
	Note    Note
	Content string
	To      string
}

type VideoNote struct {
	Note    Note
	Content string
}

type AudioNote struct {
	Note    Note
	Content string
}

type PdfNote struct {
	Note    Note
	Content string
}

type ExcalidrawNote struct {
	Note    Note
	Content string
}

type FileNote struct {
	Note    Note
	Content string
}

type CodeNote struct {
	Note     Note
	Language string
	Filename string
	Content  string
}

type QuoteNote struct {
	Note    Note
	Content string
	By      string
}

type LatexNote struct {
	Note    Note
	Content string
}
