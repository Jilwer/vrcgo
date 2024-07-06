package objects

import "time"

type announcement struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type downloadUrls struct {
	Sdk2        string `json:"sdk2"`
	Sdk3Avatars string `json:"sdk3-avatars"`
	Sdk3Worlds  string `json:"sdk3-worlds"`
	Vcc         string `json:"vcc"`
	Bootstrap   string `json:"bootstrap"`
}

type dynamicWorldRow struct {
	Index         int    `json:"index"`
	Name          string `json:"name"`
	Platform      string `json:"platform"`
	SortHeading   string `json:"sortHeading"`
	SortOrder     string `json:"sortOrder"`
	SortOwnership string `json:"sortOwnership"`
	Tag           string `json:"tag,omitempty"`
}

type event struct {
	DistanceClose             int `json:"distanceClose"`
	DistanceFactor            int `json:"distanceFactor"`
	DistanceFar               int `json:"distanceFar"`
	GroupDistance             int `json:"groupDistance"`
	MaximumBunchSize          int `json:"maximumBunchSize"`
	NotVisibleFactor          int `json:"notVisibleFactor"`
	PlayerOrderBucketSize     int `json:"playerOrderBucketSize"`
	PlayerOrderFactor         int `json:"playerOrderFactor"`
	SlowUpdateFactorThreshold int `json:"slowUpdateFactorThreshold"`
	ViewSegmentLength         int `json:"viewSegmentLength"`
}

type SystemConfig struct {
	VoiceEnableDegradation                        bool              `json:"VoiceEnableDegradation"`
	VoiceEnableReceiverLimiting                   bool              `json:"VoiceEnableReceiverLimiting"`
	Address                                       string            `json:"address"`
	Announcements                                 []announcement    `json:"announcements"`
	AppName                                       string            `json:"appName"`
	BuildVersionTag                               string            `json:"buildVersionTag"`
	ClientAPIKey                                  string            `json:"clientApiKey"`
	ClientBPSCeiling                              int               `json:"clientBPSCeiling"`
	ClientDisconnectTimeout                       int               `json:"clientDisconnectTimeout"`
	ClientReservedPlayerBPS                       int               `json:"clientReservedPlayerBPS"`
	ClientSentCountAllowance                      int               `json:"clientSentCountAllowance"`
	ContactEmail                                  string            `json:"contactEmail"`
	CopyrightEmail                                string            `json:"copyrightEmail"`
	CurrentTOSVersion                             int               `json:"currentTOSVersion"`
	DefaultAvatar                                 string            `json:"defaultAvatar"`
	DeploymentGroup                               string            `json:"deploymentGroup"`
	DevAppVersionStandalone                       string            `json:"devAppVersionStandalone"`
	DevDownloadLinkWindows                        string            `json:"devDownloadLinkWindows"`
	DevSdkURL                                     string            `json:"devSdkUrl"`
	DevSdkVersion                                 string            `json:"devSdkVersion"`
	DevServerVersionStandalone                    string            `json:"devServerVersionStandalone"`
	DisCountdown                                  time.Time         `json:"dis-countdown"`
	DisableAvatarCopying                          bool              `json:"disableAvatarCopying"`
	DisableAvatarGating                           bool              `json:"disableAvatarGating"`
	DisableCommunityLabs                          bool              `json:"disableCommunityLabs"`
	DisableCommunityLabsPromotion                 bool              `json:"disableCommunityLabsPromotion"`
	DisableEmail                                  bool              `json:"disableEmail"`
	DisableEventStream                            bool              `json:"disableEventStream"`
	DisableFeedbackGating                         bool              `json:"disableFeedbackGating"`
	DisableFrontendBuilds                         bool              `json:"disableFrontendBuilds"`
	DisableHello                                  bool              `json:"disableHello"`
	DisableOculusSubs                             bool              `json:"disableOculusSubs"`
	DisableRegistration                           bool              `json:"disableRegistration"`
	DisableSteamNetworking                        bool              `json:"disableSteamNetworking"`
	DisableTwoFactorAuth                          bool              `json:"disableTwoFactorAuth"`
	DisableUdon                                   bool              `json:"disableUdon"`
	DisableUpgradeAccount                         bool              `json:"disableUpgradeAccount"`
	DownloadLinkWindows                           string            `json:"downloadLinkWindows"`
	DownloadUrls                                  downloadUrls      `json:"downloadUrls"`
	DynamicWorldRows                              []dynamicWorldRow `json:"dynamicWorldRows"`
	Events                                        event             `json:"events"`
	GearDemoRoomID                                string            `json:"gearDemoRoomId"`
	HomeWorldID                                   string            `json:"homeWorldId"`
	HomepageRedirectTarget                        string            `json:"homepageRedirectTarget"`
	HubWorldID                                    string            `json:"hubWorldId"`
	JobsEmail                                     string            `json:"jobsEmail"`
	MessageOfTheDay                               string            `json:"messageOfTheDay"`
	ModerationEmail                               string            `json:"moderationEmail"`
	ModerationQueryPeriod                         int               `json:"moderationQueryPeriod"`
	NotAllowedToSelectAvatarInPrivateWorldMessage string            `json:"notAllowedToSelectAvatarInPrivateWorldMessage"`
	Plugin                                        string            `json:"plugin"`
	ReleaseAppVersionStandalone                   string            `json:"releaseAppVersionStandalone"`
	ReleaseSdkURL                                 string            `json:"releaseSdkUrl"`
	ReleaseSdkVersion                             string            `json:"releaseSdkVersion"`
	ReleaseServerVersionStandalone                string            `json:"releaseServerVersionStandalone"`
	SdkDeveloperFaqURL                            string            `json:"sdkDeveloperFaqUrl"`
	SdkDiscordURL                                 string            `json:"sdkDiscordUrl"`
	SdkNotAllowedToPublishMessage                 string            `json:"sdkNotAllowedToPublishMessage"`
	SdkUnityVersion                               string            `json:"sdkUnityVersion"`
	ServerName                                    string            `json:"serverName"`
	SupportEmail                                  string            `json:"supportEmail"`
	TimeOutWorldID                                string            `json:"timeOutWorldId"`
	TutorialWorldID                               string            `json:"tutorialWorldId"`
	UpdateRateMsMaximum                           int               `json:"updateRateMsMaximum"`
	UpdateRateMsMinimum                           int               `json:"updateRateMsMinimum"`
	UpdateRateMsNormal                            int               `json:"updateRateMsNormal"`
	UpdateRateMsUdonManual                        int               `json:"updateRateMsUdonManual"`
	UploadAnalysisPercent                         int               `json:"uploadAnalysisPercent"`
	URLList                                       []string          `json:"urlList"`
	UseReliableUDPForVoice                        bool              `json:"useReliableUdpForVoice"`
	UserUpdatePeriod                              int               `json:"userUpdatePeriod"`
	UserVerificationDelay                         int               `json:"userVerificationDelay"`
	UserVerificationRetry                         int               `json:"userVerificationRetry"`
	UserVerificationTimeout                       int               `json:"userVerificationTimeout"`
	ViveWindowsURL                                string            `json:"viveWindowsUrl"`
	WhiteListedAssetUrls                          []string          `json:"whiteListedAssetUrls"`
	WorldUpdatePeriod                             int               `json:"worldUpdatePeriod"`
	PlayerURLResolverHash                         string            `json:"player-url-resolver-hash"`
	PlayerURLResolverVersion                      string            `json:"player-url-resolver-version"`
}