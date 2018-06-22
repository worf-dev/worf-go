package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"bytes"
	"fmt"
)

//Contains the URL and access token required to use the API
type API struct {
	AccessToken string
	URL         string
}

//Describes an API error
type APIError struct {
	StatusCode    int
	OriginalError error
	Response      *http.Response
	Body          []byte
	Message       string
}

//Prints the error message for an APIError instance
func (e APIError) Error() string {
	if e.OriginalError != nil {
		return e.OriginalError.Error()
	}
	return ""
}

//Returns an HTTP client for use with the API
func (api API) GetClient() *http.Client {
	return &http.Client{}
}

//Parses a JSON response into a provided data object
func (api API) ParseJSON(response *http.Response, data interface{}) error {
	b, _ := ioutil.ReadAll(response.Body)
	err := json.NewDecoder(bytes.NewReader(b)).Decode(data)
	if err != nil {
		return &APIError{Message: "Error when parsing JSON.",
			OriginalError: err,
			Response:      response,
			StatusCode:    response.StatusCode,
			Body:          b}
	}
	return err
}

//Retrieves a JSON response from the API
func (api API) GetJSON(url string, params map[string]string, data interface{}) error {
	response, err := api.Get(url, params)

	if err != nil {
		return &APIError{Message: "Error during HTTP request.",
			OriginalError: err}
	}

	return api.ParseJSON(response, data)
}

//Retrieves a response from the API
func (api API) Get(url string, params map[string]string) (resp *http.Response, err error) {
	client := api.GetClient()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest(GET, %v, nil): %v", url, err)
	}
	accessTokenHeader := fmt.Sprintf("Bearer %s", api.AccessToken)
	req.Header.Add("Authorization", accessTokenHeader)
	resp, err = client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %v", err)
	}
	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return resp, fmt.Errorf("Request failed")
	}
	return
}
