package vrcapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type SearchUsersResp []struct {
	Bio                            string   `json:"bio"`
	BioLinks                       []string `json:"bioLinks"`
	CurrentAvatarImageURL          string   `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatarTags              []string `json:"currentAvatarTags"`
	DeveloperType                  string   `json:"developerType"`
	DisplayName                    string   `json:"displayName"`
	FallbackAvatar                 string   `json:"fallbackAvatar"`
	ID                             string   `json:"id"`
	IsFriend                       bool     `json:"isFriend"`
	LastPlatform                   string   `json:"last_platform"`
	ProfilePicOverride             string   `json:"profilePicOverride"`
	Pronouns                       string   `json:"pronouns"`
	Status                         string   `json:"status"`
	StatusDescription              string   `json:"statusDescription"`
	Tags                           []string `json:"tags"`
	UserIcon                       string   `json:"userIcon"`
	Location                       string   `json:"location"`
	FriendKey                      string   `json:"friendKey"`
}

// SearchUsers returns a list of users based on a text query.
func (c *VRCApiClient) SearchUsers(searchQuery string) (SearchUsersResp, error) {
	u := c.BaseURL.String() + "/users"

	q := url.Values{}
	q.Add("search", searchQuery)
	u += "?" + q.Encode()

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return SearchUsersResp{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.AddCookie(&http.Cookie{Name: "auth", Value: c.AuthCookie})
	req.AddCookie(&http.Cookie{Name: "twoFactorAuth", Value: c.TwoFactorAuthCookie})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return SearchUsersResp{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return SearchUsersResp{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return SearchUsersResp{}, errors.New("API returned non-200 status code: " + resp.Status)
	}

	var users SearchUsersResp
	err = json.Unmarshal(body, &users)
	if err != nil {
		return SearchUsersResp{}, err
	}

	return users, nil
}

type GetUserByIDResp struct {
	AllowAvatarCopying bool `json:"allowAvatarCopying"`
	Badges             []struct {
		AssignedAt       time.Time `json:"assignedAt"`
		BadgeDescription string    `json:"badgeDescription"`
		BadgeID          string    `json:"badgeId"`
		BadgeImageURL    string    `json:"badgeImageUrl"`
		BadgeName        string    `json:"badgeName"`
		Hidden           bool      `json:"hidden"`
		Showcased        bool      `json:"showcased"`
		UpdatedAt        time.Time `json:"updatedAt"`
	} `json:"badges"`
	Bio                            string   `json:"bio"`
	BioLinks                       []string `json:"bioLinks"`
	CurrentAvatarImageURL          string   `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatarTags              []string `json:"currentAvatarTags"`
	DateJoined                     string   `json:"date_joined"`
	DeveloperType                  string   `json:"developerType"`
	DisplayName                    string   `json:"displayName"`
	FriendKey                      string   `json:"friendKey"`
	FriendRequestStatus            string   `json:"friendRequestStatus"`
	ID                             string   `json:"id"`
	InstanceID                     string   `json:"instanceId"`
	IsFriend                       bool     `json:"isFriend"`
	LastActivity                   string   `json:"last_activity"`
	LastLogin                      string   `json:"last_login"`
	LastPlatform                   string   `json:"last_platform"`
	Location                       string   `json:"location"`
	Note                           string   `json:"note"`
	ProfilePicOverride             string   `json:"profilePicOverride"`
	ProfilePicOverrideThumbnail    string   `json:"profilePicOverrideThumbnail"`
	Pronouns                       string   `json:"pronouns"`
	State                          string   `json:"state"`
	Status                         string   `json:"status"`
	StatusDescription              string   `json:"statusDescription"`
	Tags                           []string `json:"tags"`
	TravelingToInstance            string   `json:"travelingToInstance"`
	TravelingToLocation            string   `json:"travelingToLocation"`
	TravelingToWorld               string   `json:"travelingToWorld"`
	UserIcon                       string   `json:"userIcon"`
	WorldID                        string   `json:"worldId"`
}

// GetUserByID returns user information about a specific user using their ID.
func (c *VRCApiClient) GetUserByID(userID string) (GetUserByIDResp, error) {
	u := c.BaseURL.String() + "/users/" + userID

	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return GetUserByIDResp{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", c.UserAgent)
	req.AddCookie(&http.Cookie{Name: "auth", Value: c.AuthCookie})
	req.AddCookie(&http.Cookie{Name: "twoFactorAuth", Value: c.TwoFactorAuthCookie})

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return GetUserByIDResp{}, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return GetUserByIDResp{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return GetUserByIDResp{}, errors.New("API returned non-200 status code: " + resp.Status)
	}

	var user GetUserByIDResp
	err = json.Unmarshal(body, &user)
	if err != nil {
		return GetUserByIDResp{}, err
	}

	return user, nil
}
