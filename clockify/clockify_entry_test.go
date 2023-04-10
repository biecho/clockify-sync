package clockify

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDeserializeClockifyEntry(t *testing.T) {
	jsonStr := `{
	  "id": "642c29b43c8e1d287f9af847",
	  "billable": false,
	  "description": "YT-725",
	  "hourlyRate": {
		"amount": 0,
		"currency": "USD"
	  },
	  "isLocked": false,
	  "project": {
		"id": "5f1af2c6f5c6fb0046e406f8",
		"name": "7007 - Code Review",
		"hourlyRate": {
		  "amount": 0,
		  "currency": ""
		},
		"clientId": "5f1aef91f5c6fb1a46e3fdf2",
		"workspaceId": "5f1aef1d41695b43ce6b5ae1",
		"billable": false,
		"memberships": [
		  {
			"hourlyRate": {
			  "amount": 0,
			  "currency": ""
			},
			"membershipStatus": "ACTIVE",
			"membershipType": "PROJECT",
			"target": "",
			"userId": "5cd95e50f15c98690babecbc"
		  }
		],
		"color": "#00BCD4",
		"timeEstimate": {
		  "type": "",
		  "active": false,
		  "resetOptions": null,
		  "estimate": ""
		},
		"budgetEstimate": {
		  "type": "",
		  "active": false,
		  "resetOptions": null,
		  "estimate": 0
		},
		"archived": false,
		"duration": "",
		"clientName": "Product Development",
		"note": "",
		"template": false,
		"public": true
	  },
	  "projectId": "5f1af2c6f5c6fb0046e406f8",
	  "tags": [],
	  "task": null,
	  "timeInterval": {
		"duration": "PT24M31S",
		"end": "2023-04-04T14:08:50Z",
		"start": "2023-04-04T13:44:19Z"
	  },
	  "totalBillable": 0,
	  "user": null,
	  "workspaceId": "5f1aef1d41695b43ce6b5ae1"
	}`

	expected := Entry{
		ID:          "642c29b43c8e1d287f9af847",
		Billable:    false,
		Description: "YT-725",
		HourlyRate: &HourlyRate{
			Amount:   0,
			Currency: "USD",
		},
		IsLocked: false,
		Project: &Project{
			ID:   "5f1af2c6f5c6fb0046e406f8",
			Name: "7007 - Code Review",
			HourlyRate: HourlyRate{
				Amount:   0,
				Currency: "",
			},
			ClientID:    "5f1aef91f5c6fb1a46e3fdf2",
			WorkspaceID: "5f1aef1d41695b43ce6b5ae1",
			Billable:    false,
			Memberships: []Membership{
				{
					HourlyRate: &HourlyRate{
						Amount:   0,
						Currency: "",
					},
					MembershipStatus: "ACTIVE",
					MembershipType:   "PROJECT",
					Target:           "",
					UserId:           "5cd95e50f15c98690babecbc",
				},
			},
			Color: "#00BCD4",
			TimeEstimate: TimeEstimate{
				Type:     "",
				Active:   false,
				Estimate: "",
			},
			BudgetEstimate: BudgetEstimate{
				Type:     "",
				Active:   false,
				Estimate: 0,
			},
			Archived:   false,
			Duration:   "",
			ClientName: "Product Development",
			Note:       "",
			Template:   false,
			Public:     true,
		},
		ProjectID: "5f1af2c6f5c6fb0046e406f8",
		Tags:      []string{},
		TimeInterval: &TimeInterval{
			Start:    time.Date(2023, 4, 4, 13, 44, 19, 0, time.UTC),
			End:      time.Date(2023, 4, 4, 14, 8, 50, 0, time.UTC),
			Duration: "PT24M31S",
		},
		TotalBillable: 0,
		User:          nil,
		WorkspaceID:   "5f1aef1d41695b43ce6b5ae1",
	}

	var actual Entry
	err := json.Unmarshal([]byte(jsonStr), &actual)
	require.NoError(t, err)

	assert.Equal(t, expected, actual)
}
