package vrcapi

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/Jilwer/vrcgo/vrcapi/objects"
	"io"
	"net/http"
	"net/url"
	"reflect"
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

// GetUserGroupRequests returns the groups a specific user has requested to be invited in using their ID.
func (c *VRCApiClient) GetUserGroupRequests(userID string) ([]objects.Group, error) {
	u := c.BaseURL.String() + "/users/" + userID + "/groups/requested"

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return []objects.Group{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.AddCookie(&http.Cookie{Name: "auth", Value: c.AuthCookie})
	req.AddCookie(&http.Cookie{Name: "twoFactorAuth", Value: c.TwoFactorAuthCookie})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return []objects.Group{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []objects.Group{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return []objects.Group{}, errors.New("API returned non-200 status code: " + resp.Status)
	}

	var groups []objects.Group
	err = json.Unmarshal(body, &groups)
	if err != nil {
		return []objects.Group{}, err
	}

	return groups, nil
}

// GetUserCurrentRepresentedGroup returns the group that the user is currently representing using their ID.
func (c *VRCApiClient) GetUserCurrentRepresentedGroup(userID string) (objects.RepresentedGroup, error) {
	u := c.BaseURL.String() + "/users/" + userID + "/groups/represented"

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return objects.RepresentedGroup{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.AddCookie(&http.Cookie{Name: "auth", Value: c.AuthCookie})
	req.AddCookie(&http.Cookie{Name: "twoFactorAuth", Value: c.TwoFactorAuthCookie})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return objects.RepresentedGroup{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return objects.RepresentedGroup{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return objects.RepresentedGroup{}, errors.New("API returned non-200 status code: " + resp.Status)
	}

	var group objects.RepresentedGroup
	err = json.Unmarshal(body, &group)
	if err != nil {
		return objects.RepresentedGroup{}, err
	}

	return group, nil
}

// UpdateUserInfo updates the information of a user with the given userID.
func (c *VRCApiClient) UpdateUserInfo(userID string, userInfo objects.UpdateUserInfoRequest) (objects.User, error) {
	u := c.BaseURL.String() + "/users/" + userID

	nonEmptyFields := make(map[string]interface{})

	val := reflect.ValueOf(userInfo)
	typ := reflect.TypeOf(userInfo)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("json")

		if !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface()) {
			nonEmptyFields[tag] = field.Interface()
		}
	}

	userInfoJson, err := json.Marshal(nonEmptyFields)
	if err != nil {
		return objects.User{}, err
	}

	req, err := http.NewRequest("PUT", u, bytes.NewBuffer(userInfoJson))
	if err != nil {
		return objects.User{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
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

	var updatedUserInfo objects.User
	err = json.Unmarshal(body, &updatedUserInfo)
	if err != nil {
		return objects.User{}, err
	}

	return updatedUserInfo, nil
}
