package notes

type CodeNote struct {
	Metadata NoteMetadata
	Language string
	Filename string
}

func (n *CodeNote) Update() {}
