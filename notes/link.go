package notes

type LinkNote struct {
	Metadata NoteMetadata
	To       string
}

func (n *LinkNote) Update() {}
