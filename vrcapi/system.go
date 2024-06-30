package vrcapi

import (
	"errors"
	"io"
	"net/http"
)

// OnlineUsers returns the number of online users in the VRChat API.
func (c *VRCApiClient) OnlineUsers() (string, error) {
	u := c.BaseURL.String() + "/visits"
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("API returned non-200 status code: " + resp.Status)
	}

	return string(body), nil

}
