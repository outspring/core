package notes

type QuoteNote struct {
	Metadata NoteMetadata
	By       string
}

func (n *QuoteNote) Update() {}
