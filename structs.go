package worf

import (
	"gitlab.com/worf/go-worf/api"
	"encoding/hex"
	"strings"
	"time"
)

//A Worf API client. Contains all routines from the generic API client.
type Client struct {
	api.API
}

type WithUUID struct {
	ID            string  `json:"id"`
}

func (w *WithUUID) BinaryID() []byte {
	rs := strings.Replace(w.ID, "-", "", 0)
	be, err := hex.DecodeString(rs)
	if err != nil {
		return nil
	}
	return be
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
	WithUUID
	Login         string  `json:"login"`
	Disabled      bool    `json:"disabled"`
	EMail         string  `json:"email"`
	NewEMail      *string `json:"new_email"`
	EMailVerified bool    `json:"email_verified"`
	SuperUser     bool    `json:"superuser"`
}

//Represents the organization of a user
type Organization struct {
	WithUUID
	Active      bool   `json:"active"`
	Description string `json:"description"`
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
