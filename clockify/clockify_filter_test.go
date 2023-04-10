package clockify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilterEntries(t *testing.T) {
	entry1 := Entry{
		ID:   "1",
		Tags: []string{"tag1", "tag2"},
	}
	entry2 := Entry{
		ID:   "2",
		Tags: []string{"tag2", "tag3"},
	}
	entry3 := Entry{
		ID:   "3",
		Tags: []string{"tag4", "tag5"},
	}
	entry4 := Entry{
		ID:   "4",
		Tags: []string{"tag5", "tag6"},
	}

	entries := []Entry{entry1, entry2, entry3, entry4}

	// Filter with tag "tag2"
	filteredEntries := FilterEntries(entries, []string{"tag2"})
	assert.Equal(t, []Entry{entry3, entry4}, filteredEntries)

	// Filter with tag "tag7"
	filteredEntries = FilterEntries(entries, []string{"tag7"})
	assert.Equal(t, entries, filteredEntries)

	// Filter with tags "tag5" and "tag6"
	filteredEntries = FilterEntries(entries, []string{"tag5", "tag6"})
	assert.Equal(t, []Entry{entry1, entry2}, filteredEntries)

	// Filter with tag ""
	filteredEntries = FilterEntries(entries, []string{""})
	assert.Equal(t, entries, filteredEntries)

	// Filter empty entries
	emptyEntries := make([]Entry, 0)
	filteredEntries = FilterEntries(emptyEntries, []string{"tag1"})
	assert.Equal(t, emptyEntries, filteredEntries)
}
