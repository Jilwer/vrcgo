package vrcapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type CommonUserFields struct {
	Bio                            string   `json:"bio"`
	BioLinks                       []string `json:"bioLinks"`
	CurrentAvatarImageURL          string   `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageUrl"`
	CurrentAvatarTags              []string `json:"currentAvatarTags"`
	DeveloperType                  string   `json:"developerType"`
	DisplayName                    string   `json:"displayName"`
	FriendKey                      string   `json:"friendKey"`
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
}

type SearchUsersResp []struct {
	CommonUserFields
	FallbackAvatar string `json:"fallbackAvatar"`
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
	CommonUserFields
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
	DateJoined                  string `json:"date_joined"`
	FriendRequestStatus         string `json:"friendRequestStatus"`
	InstanceID                  string `json:"instanceId"`
	LastActivity                string `json:"last_activity"`
	LastLogin                   string `json:"last_login"`
	Note                        string `json:"note"`
	ProfilePicOverrideThumbnail string `json:"profilePicOverrideThumbnail"`
	State                       string `json:"state"`
	TravelingToInstance         string `json:"travelingToInstance"`
	TravelingToLocation         string `json:"travelingToLocation"`
	TravelingToWorld            string `json:"travelingToWorld"`
	WorldID                     string `json:"worldId"`
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

// GetUserGroups returns user information about a specific user using their ID.
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
