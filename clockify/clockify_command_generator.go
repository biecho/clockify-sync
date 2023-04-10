package clockify

import (
	"fmt"
	"strings"
)

func CreateClockifyCommand(entry Entry, tags []string) (string, error) {
	if len(tags) == 0 {
		return "", nil
	}

	tagsArg := "--tag " + strings.Join(tags, " --tag ")
	return fmt.Sprintf("clockify-cli edit %s %s", entry.ID, tagsArg), nil
}
