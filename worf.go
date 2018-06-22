package worf

import (
	"gitlab.com/7s-worf/go-worf/api"
	"fmt"
)

//Returns a client for the Management API
func GetClient(url string, accessToken string) Client {
	return Client{API: api.API{URL: url, AccessToken: accessToken}}
}

//Returns the profile of the logged in user
func (c Client) GetUserProfile() (*UserProfile, error) {

	var userProfile *UserProfile = new(UserProfile)
	url := fmt.Sprintf("%s/user", c.URL)
	err := c.GetJSON(url, nil, userProfile)

	if err != nil {
		return nil, err
	}

	return userProfile, nil
}
