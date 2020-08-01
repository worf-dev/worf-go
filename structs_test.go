package worf

import (
	"encoding/json"
	"testing"
)

func TestProfile(t *testing.T) {
	testData := `
{
	"access_token": {
		"created_at": "2018-11-30T16:31:17.358276Z",
		"description": null,
		"id": "94dca27e-c062-45ff-bebf-7663eff0de7d",
		"last_used_at": "2018-11-30T15:31:18.029008Z",
		"last_used_from": "127.0.0.1",
		"renews_when_used": true,
		"scopes": [
		"admin"
		],
		"valid_until": null
	},
  "organizations": [
	{
		"active": true,
		"description": "test",
		"id": "14eba2d9-5b68-4b8c-b431-d900879334a9",
		"name": "Test",
		"roles": [
		  {
			"confirmed": true,
			"role": "superuser"
		  }
		]
	}
   ],
	"user": {
		"created_at": "2018-11-30T16:31:17.340251Z",
		"disabled": false,
		"display_name": null,
		"email": "user@example.com",
		"id": "8faf86b1-96b9-43a6-bbd5-1ca7b7cf7dda",
		"language": "en",
		"new_email": null,
		"superuser": false,
		"updated_at": "2018-11-30T16:31:17.340251Z"
	}
}
`
	var profile UserProfile
	err := json.Unmarshal([]byte(testData), &profile)
	if err != nil {
		t.Error(err)
	}
	if len(profile.Organizations) == 0 {
		t.Errorf("Organizations length should not be 0!")
	}
	pe := profile.Organizations[0]
	if pe.Active != true ||
		pe.Description != "test" ||
		pe.ID != "14eba2d9-5b68-4b8c-b431-d900879334a9" ||
		pe.Name != "Test" ||
		pe.Roles[0].Confirmed != true ||
		pe.Roles[0].Role != "superuser" {
		t.Errorf("Struct content doesn't match")
	}
}

func TestEmptyProfile(t *testing.T) {
	testData := `
	{
		"access_token": {
		  "created_at": "2018-11-30T16:31:17.358276Z",
		  "description": null,
		  "id": "94dca27e-c062-45ff-bebf-7663eff0de7d",
		  "last_used_at": "2018-11-30T15:31:18.029008Z",
		  "last_used_from": "127.0.0.1",
		  "renews_when_used": true,
		  "scopes": [
			"admin"
		  ],
		  "valid_until": null
		},
		"primary_organization": null,
		"user": {
		  "created_at": "2018-11-30T16:31:17.340251Z",
		  "disabled": false,
		  "display_name": null,
		  "email": "user@example.com",
		  "id": "8faf86b1-96b9-43a6-bbd5-1ca7b7cf7dda",
		  "language": "en",
		  "new_email": null,
		  "superuser": false,
		  "updated_at": "2018-11-30T16:31:17.340251Z"
		}
	  }
	`
	var profile UserProfile
	err := json.Unmarshal([]byte(testData), &profile)
	if err != nil {
		t.Error(err)
	}
	if len(profile.Organizations) != 0 {
		t.Errorf("Organizations length should be 0!")
	}
}
