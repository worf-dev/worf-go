package worf

import (
	"gitlab.com/kj7s/go-helpers/api"
	"time"
)

//A Worf API client. Contains all routines from the generic API client.
type Client struct {
	api.API
}

//Represents an error message returned by the API.
type APIErrorMessage struct {
	Message string `json:"message"`
}

//Represents an access token returned by the API.
type AccessToken struct {
	CreatedAt      *time.Time `json:"created_at"`
	ValidUntil     *time.Time `json:"valid_until"`
	LastUsedAt      time.Time `json:"last_used_at"`
	RenewsWhenUsed bool       `json:"renews_when_used"`
	Scopes         []string   `json:"scopes"`
}

//Represents a user returned by the API.
type User struct {
	Login          string       `json:"login"`
	Disabled       bool         `json:"disabled"`
	EMail          string       `json:"email"`
	NewEMail       *string      `json:"new_email"`
	EMailVerified  bool         `json:"email_verified"`
	ID             string       `json:"id"`
	SuperUser      bool         `json:"superuser"`
}

//Represents a user profile returned by the API.
type UserProfile struct {
	User        User        `json:"user"`
	AccessToken AccessToken `json:"access_token"`
}
