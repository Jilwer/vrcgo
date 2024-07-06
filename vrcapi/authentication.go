package vrcapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
CHECK USER EXISTS
*/

const (
	FilterEmail       = "email"
	FilterUsername    = "username"
	FilterDisplayName = "displayName"
)

type CheckUserExistsResp struct {
	UserExists bool `json:"userExists"`
}

// CheckUserExists checks if a user exists in the VRChat API.
func (c *VRCApiClient) CheckUserExists(filter, query string) (bool, error) {

	url := fmt.Sprintf("%s/exists?%s=%s", AuthURL, filter, query)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return false, fmt.Errorf("API returned non-200 status code: %s", resp.Status)
	}

	var userExists CheckUserExistsResp
	err = json.NewDecoder(resp.Body).Decode(&userExists)
	if err != nil {
		return false, err

	}

	return userExists.UserExists, nil
}

/*
GET CURRENT USER
*/

// TODO
// GetCurrentUser returns the current user info in the VRChat API. Also user for logging in.
func (c *VRCApiClient) GetCurrentUser() {

}

// TODO
// Verifies the email OTP in the VRChat API and sets the TwoFactorAuthCookie.
func (c *VRCApiClient) VerifyEmailOTP() {

}

// TODO
// Verifies the TwoFactorAuthCookie in the VRChat API.
func VerifyAuthToken() {

}

// TODO
// Invalidates the login session in the VRChat API.
func (c *VRCApiClient) Logout() {

}
