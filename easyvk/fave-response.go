package easyvk

// A FaveUsers describes a list of users
// whom the current user has bookmarked.
type FaveUsers struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID int `json:"id"`
			FirstName string `json:"first_name"`
			LastName string `json:"last_name"`
		} `json:"items"`
	} `json:"response"`
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

// A FavePhotos describes a list of photos
// that the current user has bookmarked.
type FavePhotos struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID int `json:"id"`
			AlbumID int `json:"album_id"`
			OwnerID int `json:"owner_id"`
			UserID int `json:"user_id"`
			Photo75 string `json:"photo_75"`
			Photo130 string `json:"photo_130"`
			Photo604 string `json:"photo_604"`
			Photo807 string `json:"photo_807"`
			Photo1280 string `json:"photo_1280"`
			Photo2560 string `json:"photo_2560"`
			Width int `json:"width"`
			Height int `json:"height"`
			Text string `json:"text"`
			Date int `json:"date"`
			PostID int `json:"post_id"`
			AccessKey string `json:"access_key"`
		} `json:"items"`
	} `json:"response"`
}