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

// A FavePhotos describes a slice of photos
// that the current user has bookmarked.
type FavePhotos []struct {
	AccessKey string `json:"access_key"`
	Aid int `json:"aid"`
	Created int `json:"created"`
	Height int `json:"height"`
	OwnerID int `json:"owner_id"`
	Pid int `json:"pid"`
	PostID int `json:"post_id,omitempty"`
	Src string `json:"src"`
	SrcBig string `json:"src_big"`
	SrcSmall string `json:"src_small"`
	SrcXbig string `json:"src_xbig"`
	SrcXxbig string `json:"src_xxbig"`
	SrcXxxbig string `json:"src_xxxbig,omitempty"`
	Text string `json:"text"`
	UserID int `json:"user_id"`
	Width int `json:"width"`
}