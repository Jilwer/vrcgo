package vrcapi

import (
	"encoding/json"
	"errors"
	"github.com/Jilwer/vrcgo/vrcapi/objects"
	"io"
	"net/http"
	"net/url"
)

// SearchUsers returns a list of users based on a text query.
func (c *VRCApiClient) SearchUsers(searchQuery string) ([]objects.LimitedUser, error) {
	u := c.BaseURL.String() + "/users"

	q := url.Values{}
	q.Add("search", searchQuery)
	u += "?" + q.Encode()

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return []objects.LimitedUser{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.AddCookie(&http.Cookie{Name: "auth", Value: c.AuthCookie})
	req.AddCookie(&http.Cookie{Name: "twoFactorAuth", Value: c.TwoFactorAuthCookie})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []objects.LimitedUser{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []objects.LimitedUser{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []objects.LimitedUser{}, errors.New("API returned non-200 status code: " + resp.Status)
	}

	var users []objects.LimitedUser
	err = json.Unmarshal(body, &users)
	if err != nil {
		return []objects.LimitedUser{}, err
	}

	return users, nil
}

// GetUserByID returns user information about a specific user using their ID.
func (c *VRCApiClient) GetUserByID(userID string) (objects.User, error) {
	u := c.BaseURL.String() + "/users/" + userID

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return objects.User{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.AddCookie(&http.Cookie{Name: "auth", Value: c.AuthCookie})
	req.AddCookie(&http.Cookie{Name: "twoFactorAuth", Value: c.TwoFactorAuthCookie})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return objects.User{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return objects.User{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return objects.User{}, errors.New("API returned non-200 status code: " + resp.Status)
	}

	var user objects.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return objects.User{}, err
	}

	return user, nil
}

// GetUserGroups returns the groups a specific user is in using their ID.
func (c *VRCApiClient) GetUserGroups(userID string) ([]objects.LimitedUserGroup, error) {
	u := c.BaseURL.String() + "/users/" + userID + "/groups"

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return []objects.LimitedUserGroup{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.AddCookie(&http.Cookie{Name: "auth", Value: c.AuthCookie})
	req.AddCookie(&http.Cookie{Name: "twoFactorAuth", Value: c.TwoFactorAuthCookie})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []objects.LimitedUserGroup{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []objects.LimitedUserGroup{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []objects.LimitedUserGroup{}, errors.New("API returned non-200 status code: " + resp.Status)
	}

	var groups []objects.LimitedUserGroup
	err = json.Unmarshal(body, &groups)
	if err != nil {
		return []objects.LimitedUserGroup{}, err
	}

	return groups, nil
}
