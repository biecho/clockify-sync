package youtrack

import (
	"github.com/biecho/clockify-sync/clockify"
	"regexp"
	"testing"
	"time"
)

func TestCreateYouTrackCommand_Success(t *testing.T) {
	entry := clockify.Entry{
		ID:          "642c00b3fcde1e287f9af847",
		Description: "This is a test description YT-1234",
		Project: &clockify.Project{
			Name: "7007 - Code Review",
		},
		Tags: []string{"tag1", "tag2"},
		TimeInterval: &clockify.TimeInterval{
			Start:    time.Date(2023, 4, 14, 13, 44, 19, 0, time.UTC),
			End:      time.Date(2023, 4, 14, 14, 8, 50, 0, time.UTC),
			Duration: "PT24M31S",
		},
	}

	syncCfg := SyncConfig{
		ProjectMapping: map[string]string{
			"7007 - Code Review": "Review",
		},
		IssueIDRegex: regexp.MustCompile(`YT-(\d+)`),
	}

	expectedCmd := "yt add work --issue-id YT-1234 --type Review --date 2023-04-14 --duration 24m31s"

	cmd, err := CreateYouTrackCommand(entry, syncCfg)
	if err != nil {
		t.Errorf("Failed to create command: %v", err)
		return
	}

	if cmd != expectedCmd {
		t.Errorf("Unexpected command. Expected: %s, got: %s", expectedCmd, cmd)
		return
	}
}

func TestCreateYouTrackCommand_InvalidProject(t *testing.T) {
	entry := clockify.Entry{
		ID:      "642c00b3fcde1e287f9af847",
		Project: nil,
		TimeInterval: &clockify.TimeInterval{
			Start:    time.Date(2023, 4, 14, 13, 44, 19, 0, time.UTC),
			End:      time.Date(2023, 4, 14, 14, 8, 50, 0, time.UTC),
			Duration: "PT24M31S",
		},
	}

	syncCfg := SyncConfig{
		ProjectMapping: map[string]string{
			"7007 - Code Review": "Review",
		},
		IssueIDRegex: regexp.MustCompile(`YT-(\d+)`),
	}

	_, err := CreateYouTrackCommand(entry, syncCfg)
	if err == nil {
		t.Errorf("Expected error, but got none")
		return
	}

	expectedErrMsg := "project is nil in entry '642c00b3fcde1e287f9af847'"
	if err.Error() != expectedErrMsg {
		t.Errorf("Unexpected error message. Expected: %s, got: %v", expectedErrMsg, err)
		return
	}
}

func TestCreateYouTrackCommand_InvalidTimeInterval(t *testing.T) {
	entry := clockify.Entry{
		ID:          "642c00b3fcde1e287f9af847",
		Description: "This is a test description YT-1234",
		Project: &clockify.Project{
			Name: "7007 - Code Review",
		},
		Tags:         []string{"tag1", "tag2"},
		TimeInterval: nil,
	}

	syncCfg := SyncConfig{
		ProjectMapping: map[string]string{
			"7007 - Code Review": "Review",
		},
		IssueIDRegex: regexp.MustCompile(`YT-(\d+)`),
	}

	_, err := CreateYouTrackCommand(entry, syncCfg)
	if err == nil {
		t.Error("Expected an error due to invalid time interval")
	}
}

func TestCreateYouTrackCommand_NoMapping(t *testing.T) {
	entry := clockify.Entry{
		ID:          "642c00b3fcde1e287f9af847",
		Description: "This is a test description YT-1234",
		Project: &clockify.Project{
			Name: "invalid project name",
		},
		Tags: []string{"tag1", "tag2"},
		TimeInterval: &clockify.TimeInterval{
			Start:    time.Now(),
			End:      time.Now().Add(1 * time.Hour),
			Duration: "PT1H",
		},
	}

	syncCfg := SyncConfig{
		ProjectMapping: map[string]string{
			"7007 - Code Review": "Review",
		},
		IssueIDRegex: regexp.MustCompile(`YT-(\d+)`),
	}

	_, err := CreateYouTrackCommand(entry, syncCfg)
	if err == nil {
		t.Error("Expected an error due to no mapping found for project")
	}
}

func TestCreateYouTrackCommand_InvalidIssueID(t *testing.T) {
	entry := clockify.Entry{
		ID:          "642c00b3fcde1e287f9af847",
		Description: "This is a test description without an issue ID",
		Project: &clockify.Project{
			Name: "7007 - Code Review",
		},
		Tags: []string{"tag1", "tag2"},
		TimeInterval: &clockify.TimeInterval{
			Start:    time.Now(),
			End:      time.Now().Add(1 * time.Hour),
			Duration: "PT1H",
		},
	}

	syncCfg := SyncConfig{
		ProjectMapping: map[string]string{
			"7007 - Code Review": "Review",
		},
		IssueIDRegex: regexp.MustCompile(`YT-(\d+)`),
	}

	_, err := CreateYouTrackCommand(entry, syncCfg)
	if err == nil {
		t.Error("Expected an error due to invalid YouTrack issue ID in description")
	}
}

func TestCreateYouTrackCommand_InvalidDuration(t *testing.T) {
	entry := clockify.Entry{
		ID:          "642c00b3fcde1e287f9af847",
		Description: "This is a test description YT-1234",
		Project: &clockify.Project{
			Name: "7007 - Code Review",
		},
		Tags: []string{"tag1", "tag2"},
		TimeInterval: &clockify.TimeInterval{
			Start:    time.Date(2023, 4, 14, 13, 44, 19, 0, time.UTC),
			End:      time.Date(2023, 4, 14, 14, 8, 50, 0, time.UTC),
			Duration: "invalid-duration",
		},
	}

	syncCfg := SyncConfig{
		ProjectMapping: map[string]string{
			"7007 - Code Review": "Review",
		},
		IssueIDRegex: regexp.MustCompile(`YT-(\d+)`),
	}

	_, err := CreateYouTrackCommand(entry, syncCfg)
	if err == nil {
		t.Error("Expected an error, but got none")
		return
	}
}
