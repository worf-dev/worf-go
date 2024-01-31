package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Contains the URL and access token required to use the API
type API struct {
	AccessToken string
	URL         string
	client      *http.Client
}

// Describes an API error
type APIError struct {
	StatusCode    int
	OriginalError error
	Response      *http.Response
	Body          []byte
	Message       string
}

func MakeAPI(url, accessToken string) *API {
	return &API{
		URL:         url,
		AccessToken: accessToken,
		client:      &http.Client{},
	}
}

// Prints the error message for an APIError instance
func (e APIError) Error() string {
	if e.OriginalError != nil {
		return e.OriginalError.Error()
	}
	return ""
}

// Parses a JSON response into a provided data object
func (api *API) ParseJSON(response *http.Response, data interface{}) error {
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

// Retrieves a JSON response from the API
func (api *API) JSON(method string, url string, params map[string]string, data interface{}) error {
	response, err := api.Request(method, fmt.Sprintf("%s/%s", api.URL, url), params)

	if err != nil {
		return &APIError{Message: "Error during HTTP request.",
			OriginalError: err}
	}

	if data == nil {
		// we don't need to parse a response
		return nil
	}

	return api.ParseJSON(response, data)
}

// Retrieves a response from the API
func (api *API) Request(method string, url string, params map[string]string) (resp *http.Response, err error) {

	var data io.Reader

	// we always add parameters as a JSON body
	if params != nil {

		jsonData, err := json.Marshal(params)

		if err != nil {
			return nil, err
		}

		data = bytes.NewBuffer(jsonData)
	}

	req, err := http.NewRequest(method, url, data)

	if data != nil {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}

	if err != nil {
		return nil, fmt.Errorf("http.NewRequest(%s, %v, nil): %v", method, url, err)
	}

	if api.AccessToken != "" {
		accessTokenHeader := fmt.Sprintf("Bearer %s", api.AccessToken)
		req.Header.Add("Authorization", accessTokenHeader)
	}

	resp, err = api.client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("client.Do: %v", err)
	}

	if resp.StatusCode >= 300 || resp.StatusCode < 200 {
		return resp, fmt.Errorf("Request failed")
	}

	return
}
