package youtrack

import (
	"fmt"
	iso8601 "github.com/ChannelMeter/iso8601duration"
	"github.com/biecho/clockify-sync/clockify"
	"regexp"
	"time"
)

type SyncConfig struct {
	ProjectMapping map[string]string
	IssueIDRegex   *regexp.Regexp
}

func CreateYouTrackCommand(entry clockify.Entry, config SyncConfig) (string, error) {
	if err := validateEntry(entry); err != nil {
		return "", err
	}

	workItemType, err := getWorkItemType(entry.Project.Name, config.ProjectMapping)
	if err != nil {
		return "", err
	}

	issueID, err := getIssueID(entry.Description, config.IssueIDRegex)
	if err != nil {
		return "", err
	}

	duration, err := parseDuration(entry.TimeInterval.Duration, entry.ID)
	if err != nil {
		return "", err
	}

	command := fmt.Sprintf(
		"yt add work --issue-id %s --type %s --date %s --duration %s",
		issueID,
		workItemType,
		entry.TimeInterval.Start.Format("2006-01-02"),
		duration.String(),
	)

	return command, nil
}

func validateEntry(entry clockify.Entry) error {
	if entry.Project == nil {
		return fmt.Errorf("project is nil in entry '%s'", entry.ID)
	}

	if entry.TimeInterval == nil {
		return fmt.Errorf("time interval is nil in entry '%s'", entry.ID)
	}

	return nil
}

func getWorkItemType(projectName string, projectMapping map[string]string) (string, error) {
	workItemType, ok := projectMapping[projectName]
	if !ok {
		return "", fmt.Errorf("no mapping found for project '%s'", projectName)
	}

	return workItemType, nil
}

func getIssueID(description string, issueIDRegex *regexp.Regexp) (string, error) {
	issueID := issueIDRegex.FindString(description)
	if issueID == "" {
		return "", fmt.Errorf("unable to find YouTrack issue ID in description '%s'", description)
	}
	return issueID, nil
}

func parseDuration(duration string, entryID string) (time.Duration, error) {
	iso8601Duration, err := iso8601.FromString(duration)
	if err != nil {
		return 0, fmt.Errorf("unable to parse duration '%s' in entry '%s': %v", duration, entryID, err)
	}

	return iso8601Duration.ToDuration(), nil
}
