package worf

import (
	"gitlab.com/worf/go-worf/api"
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
	LastUsedAt     time.Time  `json:"last_used_at"`
	RenewsWhenUsed bool       `json:"renews_when_used"`
	Scopes         []string   `json:"scopes"`
}

//Represents a user returned by the API.
type User struct {
	Login         string  `json:"login"`
	Disabled      bool    `json:"disabled"`
	EMail         string  `json:"email"`
	NewEMail      *string `json:"new_email"`
	EMailVerified bool    `json:"email_verified"`
	ID            string  `json:"id"`
	SuperUser     bool    `json:"superuser"`
}

//Represents the organization of a user
type Organization struct {
	Active      bool   `json:"active"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Name        string `json:"name"`
	Roles       []Role `json:"roles"`

}

//Represents the role of a user in an organization
type Role struct {
	Confirmed bool   `json:"confirmed"`
	Role      string `json:"role"`
}

//Represents a user profile returned by the API.
type UserProfile struct {
	User                User          `json:"user"`
	AccessToken         AccessToken   `json:"access_token"`
	Organization        *Organization `json:"organization,omitempty"`
}
