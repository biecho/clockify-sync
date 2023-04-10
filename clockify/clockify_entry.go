package clockify

import "time"

type Entry struct {
	ID            string        `json:"id,omitempty"`
	Billable      bool          `json:"billable,omitempty"`
	Description   string        `json:"description,omitempty"`
	HourlyRate    *HourlyRate   `json:"hourlyRate,omitempty"`
	IsLocked      bool          `json:"isLocked,omitempty"`
	Project       *Project      `json:"project,omitempty"`
	ProjectID     string        `json:"projectId,omitempty"`
	Tags          []string      `json:"tags,omitempty"`
	TimeInterval  *TimeInterval `json:"timeInterval,omitempty"`
	TotalBillable int64         `json:"totalBillable,omitempty"`
	User          *User         `json:"user,omitempty"`
	WorkspaceID   string        `json:"workspaceId,omitempty"`
}

type HourlyRate struct {
	Amount   float64 `json:"amount,omitempty"`
	Currency string  `json:"currency,omitempty"`
}

type Project struct {
	ID             string         `json:"id,omitempty"`
	Name           string         `json:"name,omitempty"`
	HourlyRate     HourlyRate     `json:"hourlyRate,omitempty"`
	ClientID       string         `json:"clientId,omitempty"`
	WorkspaceID    string         `json:"workspaceId,omitempty"`
	Billable       bool           `json:"billable,omitempty"`
	Memberships    []Membership   `json:"memberships,omitempty"`
	Color          string         `json:"color,omitempty"`
	TimeEstimate   TimeEstimate   `json:"timeEstimate,omitempty"`
	BudgetEstimate BudgetEstimate `json:"budgetEstimate,omitempty"`
	Archived       bool           `json:"archived,omitempty"`
	Duration       string         `json:"duration,omitempty"`
	ClientName     string         `json:"clientName,omitempty"`
	Note           string         `json:"note,omitempty"`
	Template       bool           `json:"template,omitempty"`
	Public         bool           `json:"public,omitempty"`
}

type TimeEstimate struct {
	Type        string `json:"type,omitempty"`
	Active      bool   `json:"active,omitempty"`
	ResetOption string `json:"resetOptions,omitempty"`
	Estimate    string `json:"estimate,omitempty"`
}

type BudgetEstimate struct {
	Type        string `json:"type,omitempty"`
	Active      bool   `json:"active,omitempty"`
	ResetOption string `json:"resetOptions,omitempty"`
	Estimate    int64  `json:"estimate,omitempty"`
}

type Membership struct {
	HourlyRate       *HourlyRate `json:"hourlyRate,omitempty"`
	MembershipStatus string      `json:"membershipStatus,omitempty"`
	MembershipType   string      `json:"membershipType,omitempty"`
	Target           string      `json:"target,omitempty"`
	UserId           string      `json:"userId,omitempty"`
}

type TimeInterval struct {
	Start    time.Time `json:"start,omitempty"`
	End      time.Time `json:"end,omitempty"`
	Duration string    `json:"duration,omitempty"`
}

type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Active   bool   `json:"active,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}
