package easyvk

// A FaveUsers describes a slice of users
// whom the current user has bookmarked.
type FaveUsers []struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UID int `json:"uid"`
}
