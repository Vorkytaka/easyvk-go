package easyvk

// A FaveUsers describes a slice of users
// whom the current user has bookmarked.
type FaveUsers []struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	UID int `json:"uid"`
}

// A FaveLinks describes a slice of links
// that the current user has bookmarked.
type FaveLinks []struct {
	ID string `json:"id"`
	URL string `json:"url"`
	Title string `json:"title"`
	Description string `json:"description"`
	ImageSrc string `json:"image_src"`
	ImageMiddle string `json:"image_middle"`
	ImageBig string `json:"image_big"`
}