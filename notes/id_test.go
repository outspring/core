package notes

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewNoteId(t *testing.T) {
	tests := []struct {
		time     time.Time
		expected string
	}{
		{time: time.Unix(973940400, 0).UTC(), expected: "20001111110000"}, // My birthday :)
		{time: time.Unix(0, 0).UTC(), expected: "19700101000000"},         // Epoch time
	}

	for _, test := range tests {
		id := NewNoteId(test.time)

		assert.Equal(t, test.expected, id.value)
	}
}
