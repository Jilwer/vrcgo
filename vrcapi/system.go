package vrcapi

import (
	"encoding/json"
	"errors"
	"github.com/Jilwer/vrcgo/vrcapi/objects"
	"io"
	"net/http"
)

// GetOnlineUsers returns the number of online users in the VRChat API.
func (c *VRCApiClient) GetOnlineUsers() (string, error) {
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

// GetSystemTime returns the current time of the VRChat API.
func (c *VRCApiClient) GetSystemTime() (string, error) {
	u := c.BaseURL.String() + "/time"
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

// GetSystemConfig Returns the system configuration of the VRChat API as a SystemConfigResp struct.
func (c *VRCApiClient) GetSystemConfig() (objects.SystemConfig, error) {
	u := c.BaseURL.String() + "/config"
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return objects.SystemConfig{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return objects.SystemConfig{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return objects.SystemConfig{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return objects.SystemConfig{}, errors.New("API returned non-200 status code: " + resp.Status)
	}

	var config objects.SystemConfig
	err = json.Unmarshal(body, &config)
	if err != nil {
		return objects.SystemConfig{}, err
	}

	return config, nil
}
