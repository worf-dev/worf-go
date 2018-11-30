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
		"primary_organization": {
		  "active": true,
		  "description": "test",
		  "id": "566807f1-699f-4b6e-8f1f-f65563c73f11",
		  "name": "Test",
		  "role": {
			"confirmed": true,
			"name": "superuser",
			"primary": true
		  }
		},
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
	if profile.PrimaryOrganization == nil {
		t.Errorf("Organization should not be nil!")
	}
	pe := profile.PrimaryOrganization
	if pe.Active != true ||
		 pe.Description != "test" ||
		 pe.ID != "566807f1-699f-4b6e-8f1f-f65563c73f11" ||
		 pe.Name != "Test" ||
		 pe.Role.Confirmed != true ||
		 pe.Role.Name != "superuser" ||
		 pe.Role.Primary != true {
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
	if profile.PrimaryOrganization != nil {
		t.Errorf("Organization should be nil!")
	}
}
