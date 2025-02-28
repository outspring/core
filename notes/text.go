package notes

type TextNote struct {
	Metadata NoteMetadata
}

func (n *TextNote) Update(content string) {
	n.Metadata.Content = content
}
