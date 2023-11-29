package worf

import (
	"github.com/getworf/worf-go/api"
	"net/http"
)

// Returns a client for the Management API
func MakeClient(url string, accessToken string) *Client {
	return &Client{API: api.MakeAPI(url, accessToken)}
}

// Returns the profile of the logged in user
func (c *Client) UserProfile() (*UserProfile, error) {
	userProfile := &UserProfile{}
	if err := c.JSON(http.MethodGet, "user", nil, userProfile); err != nil {
		return nil, err
	} else {
		return userProfile, nil
	}
}

func (c *Client) PasswordLogin(email, password string) (*UserProfile, error) {
	userProfile := &UserProfile{}
	if err := c.JSON(http.MethodPost, "login/password", map[string]string{"email": email, "password": password}, userProfile); err != nil {
		return nil, err
	} else {
		return userProfile, nil
	}
}
