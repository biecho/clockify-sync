package clockify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateClockifyCommand(t *testing.T) {
	entry1 := Entry{
		ID:   "1",
		Tags: []string{"tag1", "tag2"},
	}

	entry2 := Entry{
		ID:   "2",
		Tags: []string{"tag3", "tag4"},
	}

	testCases := []struct {
		name     string
		entry    Entry
		tags     []string
		expected string
	}{
		{
			name:     "empty tags",
			entry:    entry1,
			tags:     []string{},
			expected: "",
		},
		{
			name:     "single tag",
			entry:    entry1,
			tags:     []string{"tag1"},
			expected: "clockify-cli edit 1 --tag tag1",
		},
		{
			name:     "multiple tags",
			entry:    entry2,
			tags:     []string{"tag3", "tag4"},
			expected: "clockify-cli edit 2 --tag tag3 --tag tag4",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd, err := CreateClockifyCommand(tc.entry, tc.tags)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, cmd)
		})
	}
}
