package objects

// CheckUserExists represents the response indicating whether a user exists.
type CheckUserExists struct {
	UserExists bool `json:"userExists"`
}
