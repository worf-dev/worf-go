package worf

import (
	"fmt"
	"github.com/7scientists/worf-go/api"
)

//Returns a client for the Management API
func MakeClient(url string, accessToken string) *Client {
	return &Client{API: api.MakeAPI(url, accessToken)}
}

//Returns the profile of the logged in user
func (c *Client) UserProfile() (*UserProfile, error) {

	var userProfile *UserProfile = new(UserProfile)
	url := fmt.Sprintf("%s/user", c.URL)
	err := c.JSON(url, nil, userProfile)

	if err != nil {
		return nil, err
	}

	return userProfile, nil
}
