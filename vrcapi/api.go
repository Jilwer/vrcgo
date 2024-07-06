package vrcapi

import (
	"github.com/Jilwer/vrcgo/vrcapi/objects"
	"net/http"
	"net/url"
)

const BaseURL = "https://vrchat.com/api/1"
const BaseDevAPI = "https://dev-api.vrchat.cloud/api/1"
const BaseClientAPI = "https://api.vrchat.cloud/api/1"
const APIKey = "JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26"
const AuthURL = BaseURL + "/auth"
const AvatarsURL = BaseURL + "/avatars"
const UsersURL = BaseURL + "/users"
const WorldsURL = BaseURL + "/worlds"

func NewVRCApiClient(baseURL string, userAgent string) (*objects.VRCApiClient, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &objects.VRCApiClient{
		BaseURL:    base,
		UserAgent:  userAgent,
		HttpClient: &http.Client{},
	}, nil
}

func NewVRCApiClientWithAuth(baseURL string, userAgent string, authCookie string, twoFactorAuthCookie string) (*objects.VRCApiClient, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &objects.VRCApiClient{
		BaseURL:             base,
		UserAgent:           userAgent,
		AuthCookie:          authCookie,
		TwoFactorAuthCookie: twoFactorAuthCookie,
		HttpClient:          &http.Client{},
	}, nil
}
