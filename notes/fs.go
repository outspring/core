package notes

import (
	"errors"
	"fmt"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

var (
	ErrNoteWithoutType = errors.New("Note file does not have a note type")
)

func ReadNote(fs *afero.Fs, id NoteId) (NoteType, []byte, error) {
	filename := fmt.Sprintf("%s.yaml", id.Value())

	file, err := (*fs).Open(filename)

	if err != nil {
		return "", nil, err
	}

	defer file.Close()

	var fileData []byte

	_, err = file.Read(fileData)

	if err != nil {
		return "", nil, err
	}

	noteType, err := readNoteType(fileData)

	if err != nil {
		return "", nil, err
	}

	return noteType, fileData, nil
}

func readNoteType(fileData []byte) (NoteType, error) {
	var yamlData map[string]any

	err := yaml.Unmarshal(fileData, yamlData)

	if err != nil {
		return "", err
	}

	noteType, ok := yamlData["type"].(string)

	if !ok {
		return "", ErrNoteWithoutType
	}

	return NoteType(noteType), nil
}
