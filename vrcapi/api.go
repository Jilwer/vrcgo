package vrcapi

import (
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

type VRCApiClient struct {
	BaseURL   *url.URL
	UserAgent string

	AuthCookie          string
	TwoFactorAuthCookie string
	httpClient          *http.Client
}

func NewVRCApiClient(baseURL string, userAgent string) (*VRCApiClient, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &VRCApiClient{
		BaseURL:    base,
		UserAgent:  userAgent,
		httpClient: &http.Client{},
	}, nil
}

func NewVRCApiClientWithAuth(baseURL string, userAgent string, authCookie string, twoFactorAuthCookie string) (*VRCApiClient, error) {
	base, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	return &VRCApiClient{
		BaseURL:             base,
		UserAgent:           userAgent,
		AuthCookie:          authCookie,
		TwoFactorAuthCookie: twoFactorAuthCookie,
		httpClient:          &http.Client{},
	}, nil
}
