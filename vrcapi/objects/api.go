package objects

import (
	"net/http"
	"net/url"
)

// VRCApiClient represents a client for interacting with the VRC API.
type VRCApiClient struct {
	BaseURL             *url.URL
	UserAgent           string
	AuthCookie          string
	TwoFactorAuthCookie string
	HttpClient          *http.Client
}
