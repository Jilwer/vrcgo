package objects

import "time"

// Badge represents a user badge with its details.
type Badge struct {
	AssignedAt       time.Time `json:"assignedAt"`
	BadgeDescription string    `json:"badgeDescription"`
	BadgeID          string    `json:"badgeId"`
	BadgeImageURL    string    `json:"badgeImageUrl"`
	BadgeName        string    `json:"badgeName"`
	Hidden           bool      `json:"hidden"`
	Showcased        bool      `json:"showcased"`
	UpdatedAt        time.Time `json:"updatedAt"`
}

// LimitedUser represents a minimal user profile.
type LimitedUser struct {
	Bio                            string   `json:"bio"`
	BioLinks                       []string `json:"bioLinks"`
	CurrentAvatarImageURL          string   `json:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageURL"`
	CurrentAvatarTags              []string `json:"currentAvatarTags"`
	DeveloperType                  string   `json:"developerType"`
	DisplayName                    string   `json:"displayName"`
	FallbackAvatar                 string   `json:"fallbackAvatar"`
	ID                             string   `json:"id"`
	IsFriend                       bool     `json:"isFriend"`
	LastPlatform                   string   `json:"lastPlatform"`
	ProfilePicOverride             string   `json:"profilePicOverride"`
	Pronouns                       string   `json:"pronouns"`
	Status                         string   `json:"status"`
	StatusDescription              string   `json:"statusDescription"`
	Tags                           []string `json:"tags"`
	UserIcon                       string   `json:"userIcon"`
	Location                       string   `json:"location"`
	FriendKey                      string   `json:"friendKey"`
}

// User represents a full user profile.
type User struct {
	AllowAvatarCopying             bool     `json:"allowAvatarCopying"`
	Badges                         []Badge  `json:"badges"`
	Bio                            string   `json:"bio"`
	BioLinks                       []string `json:"bioLinks"`
	CurrentAvatarImageURL          string   `json:"currentAvatarImageURL"`
	CurrentAvatarThumbnailImageURL string   `json:"currentAvatarThumbnailImageURL"`
	CurrentAvatarTags              []string `json:"currentAvatarTags"`
	DateJoined                     string   `json:"dateJoined"`
	DeveloperType                  string   `json:"developerType"`
	DisplayName                    string   `json:"displayName"`
	FriendKey                      string   `json:"friendKey"`
	FriendRequestStatus            string   `json:"friendRequestStatus"`
	ID                             string   `json:"id"`
	InstanceID                     string   `json:"instanceId"`
	IsFriend                       bool     `json:"isFriend"`
	LastActivity                   string   `json:"lastActivity"`
	LastLogin                      string   `json:"lastLogin"`
	LastPlatform                   string   `json:"lastPlatform"`
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

// Group represents a group with its details.
type Group struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	ShortCode           string    `json:"shortCode"`
	Discriminator       string    `json:"discriminator"`
	Description         string    `json:"description"`
	IconURL             string    `json:"iconURL"`
	BannerURL           string    `json:"bannerURL"`
	Privacy             string    `json:"privacy"`
	OwnerID             string    `json:"ownerId"`
	Rules               string    `json:"rules"`
	Links               []string  `json:"links"`
	Languages           []string  `json:"languages"`
	IconID              string    `json:"iconId"`
	BannerID            string    `json:"bannerId"`
	MemberCount         int       `json:"memberCount"`
	MemberCountSyncedAt time.Time `json:"memberCountSyncedAt"`
	IsVerified          bool      `json:"isVerified"`
	JoinState           string    `json:"joinState"`
	Tags                []string  `json:"tags"`
	Galleries           []Gallery `json:"galleries"`
	CreatedAt           time.Time `json:"createdAt"`
	OnlineMemberCount   int       `json:"onlineMemberCount"`
	MembershipStatus    string    `json:"membershipStatus"`
	MyMember            Member    `json:"myMember"`
	Roles               []Role    `json:"roles"`
}

// Gallery represents a group's gallery with its details.
type Gallery struct {
	ID                   string    `json:"id"`
	Name                 string    `json:"name"`
	Description          string    `json:"description"`
	MembersOnly          bool      `json:"membersOnly"`
	RoleIdsToView        []string  `json:"roleIdsToView"`
	RoleIdsToSubmit      []string  `json:"roleIdsToSubmit"`
	RoleIdsToAutoApprove []string  `json:"roleIdsToAutoApprove"`
	RoleIdsToManage      []string  `json:"roleIdsToManage"`
	CreatedAt            time.Time `json:"createdAt"`
	UpdatedAt            time.Time `json:"updatedAt"`
}

// Member represents a group member with their details.
type Member struct {
	ID                          string    `json:"id"`
	GroupID                     string    `json:"groupId"`
	UserID                      string    `json:"userId"`
	RoleIds                     []string  `json:"roleIds"`
	ManagerNotes                string    `json:"managerNotes"`
	MembershipStatus            string    `json:"membershipStatus"`
	IsSubscribedToAnnouncements bool      `json:"isSubscribedToAnnouncements"`
	Visibility                  string    `json:"visibility"`
	IsRepresenting              bool      `json:"isRepresenting"`
	JoinedAt                    time.Time `json:"joinedAt"`
	BannedAt                    string    `json:"bannedAt"`
	Has2FA                      bool      `json:"has2FA"`
	Permissions                 []string  `json:"permissions"`
}

// Role represents a role within a group.
type Role struct {
	ID                string    `json:"id"`
	GroupID           string    `json:"groupId"`
	Name              string    `json:"name"`
	Description       string    `json:"description"`
	IsSelfAssignable  bool      `json:"isSelfAssignable"`
	Permissions       []string  `json:"permissions"`
	IsManagementRole  bool      `json:"isManagementRole"`
	RequiresTwoFactor bool      `json:"requiresTwoFactor"`
	RequiresPurchase  bool      `json:"requiresPurchase"`
	Order             int       `json:"order"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
}

// LimitedUserGroup represents a minimal group profile.
type LimitedUserGroup struct {
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

// RepresentedGroup represents a group that a user is representing.
type RepresentedGroup struct {
	Name             string `json:"name"`
	ShortCode        string `json:"shortCode"`
	Discriminator    string `json:"discriminator"`
	Description      string `json:"description"`
	IconID           string `json:"iconId"`
	IconURL          string `json:"iconUrl"`
	BannerID         string `json:"bannerId"`
	BannerURL        string `json:"bannerUrl"`
	Privacy          string `json:"privacy"`
	OwnerID          string `json:"ownerId"`
	MemberCount      int    `json:"memberCount"`
	GroupID          string `json:"groupId"`
	MemberVisibility string `json:"memberVisibility"`
	IsRepresenting   bool   `json:"isRepresenting"`
}
