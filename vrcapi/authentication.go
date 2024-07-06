package vrcapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/Jilwer/vrcgo/vrcapi/objects"
	"net/http"
)

/*
CHECK USER EXISTS
*/

const (
	FilterEmail       = "email"
	FilterUserID      = "username"
	FilterDisplayName = "displayName"
)

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

	var userExists objects.CheckUserExists
	err = json.NewDecoder(resp.Body).Decode(&userExists)
	if err != nil {
		return false, err

	}

	return userExists.UserExists, nil
}

/*
GET CURRENT USER
*/

// GetCurrentUser returns the current user info in the VRChat API. Also user for logging in.
func (c *VRCApiClient) GetCurrentUser(username, password string) (*objects.GetCurrentUserResp, error) {

	req, err := http.NewRequest("GET", AuthURL, nil)
	if err != nil {
		return &objects.GetCurrentUserResp{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	if username != "" && password != "" {
		req.Header.Set("Authorization", basicAuthEncode(username, password))
	} else if c.AuthCookie != "" && c.TwoFactorAuthCookie != "" {
		req.AddCookie(&http.Cookie{Name: "auth", Value: c.AuthCookie})
		req.AddCookie(&http.Cookie{Name: "twoFactorAuth", Value: c.TwoFactorAuthCookie})
	} else {
		return &objects.GetCurrentUserResp{}, fmt.Errorf("username and password are required")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return &objects.GetCurrentUserResp{}, err
	}

	if resp.StatusCode != http.StatusOK {
		var errResp objects.Error
		err = json.NewDecoder(resp.Body).Decode(&errResp)
		if err != nil {
			return &objects.GetCurrentUserResp{}, err
		}
		return &objects.GetCurrentUserResp{}, fmt.Errorf("API returned non-200 status code: %s", errResp.Error.Message)
	}

	defer resp.Body.Close()

	var user objects.GetCurrentUserResp
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return &objects.GetCurrentUserResp{}, err
	}

	c.AuthCookie = resp.Cookies()[0].Value

	return &user, nil
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

func basicAuthEncode(username, password string) string {
	return fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(username+":"+password)))
}
