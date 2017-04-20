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

// A FaveLinks describes a list of links
// that the current user has bookmarked.
type FaveLinks struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID string `json:"id"`
			URL string `json:"url"`
			Title string `json:"title"`
			Description string `json:"description"`
			Photo50 string `json:"photo_50"`
			Photo100 string `json:"photo_100"`
			Photo200 string `json:"photo_200"`
		} `json:"items"`
	} `json:"response"`
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

// A FaveVideos describes a list of videos
// that the current user has bookmarked.
type FaveVideos struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID int `json:"id"`
			OwnerID int `json:"owner_id"`
			Title string `json:"title"`
			Duration int `json:"duration"`
			Description string `json:"description"`
			Date int `json:"date"`
			Comments int `json:"comments"`
			Views int `json:"views"`
			Width int `json:"width"`
			Height int `json:"height"`
			Photo130 string `json:"photo_130"`
			Photo320 string `json:"photo_320"`
			Photo800 string `json:"photo_800"`
			CanAdd int `json:"can_add"`
		} `json:"items"`
	} `json:"response"`
}