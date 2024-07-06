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
func (c *VRCApiClient) SearchUsers(searchQuery string) (objects.LimitedUser, error) {
	u := c.BaseURL.String() + "/users"

	q := url.Values{}
	q.Add("search", searchQuery)
	u += "?" + q.Encode()

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return objects.LimitedUser{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.AddCookie(&http.Cookie{Name: "auth", Value: c.AuthCookie})
	req.AddCookie(&http.Cookie{Name: "twoFactorAuth", Value: c.TwoFactorAuthCookie})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return objects.LimitedUser{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return objects.LimitedUser{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return objects.LimitedUser{}, errors.New("API returned non-200 status code: " + resp.Status)
	}

	var users objects.LimitedUser
	err = json.Unmarshal(body, &users)
	if err != nil {
		return objects.LimitedUser{}, err
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

type GetUserGroupsResp []struct {
	ID                string    `json:"id"`
	Name              string    `json:"name"`
	ShortCode         string    `json:"shortCode"`
	Discriminator     string    `json:"discriminator"`
	Description       string    `json:"description"`
	IconID            string    `json:"iconId"`
	IconURL           string    `json:"iconUrl"`
	BannerID          string    `json:"bannerId"`
	BannerURL         string    `json:"bannerUrl"`
	Privacy           string    `json:"privacy"`
	LastPostCreatedAt time.Time `json:"lastPostCreatedAt"`
	OwnerID           string    `json:"ownerId"`
	MemberCount       int       `json:"memberCount"`
	GroupID           string    `json:"groupId"`
	MemberVisibility  string    `json:"memberVisibility"`
	IsRepresenting    bool      `json:"isRepresenting"`
	MutualGroup       bool      `json:"mutualGroup"`
	LastPostReadAt    time.Time `json:"lastPostReadAt"`
}

// GetUserGroups returns the groups a specific user is in using their ID.
func (c *VRCApiClient) GetUserGroups(userID string) (GetUserGroupsResp, error) {
	u := c.BaseURL.String() + "/users/" + userID + "/groups"

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return GetUserGroupsResp{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.AddCookie(&http.Cookie{Name: "auth", Value: c.AuthCookie})
	req.AddCookie(&http.Cookie{Name: "twoFactorAuth", Value: c.TwoFactorAuthCookie})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return GetUserGroupsResp{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetUserGroupsResp{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return GetUserGroupsResp{}, errors.New("API returned non-200 status code: " + resp.Status)
	}

	var groups GetUserGroupsResp
	err = json.Unmarshal(body, &groups)
	if err != nil {
		return GetUserGroupsResp{}, err
	}

	return groups, nil
}
