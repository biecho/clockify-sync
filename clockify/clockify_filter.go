package clockify

import (
	"github.com/thoas/go-funk"
)

func FilterEntries(entries []Entry, ignoredTags []string) []Entry {
	filtered := funk.Filter(entries, func(entry Entry) bool {
		return !funk.Contains(ignoredTags, func(tag string) bool {
			return funk.ContainsString(entry.Tags, tag)
		})
	}).([]Entry)

	return filtered
}
