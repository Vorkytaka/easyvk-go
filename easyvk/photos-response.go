package easyvk

// WallUploadServer describes the server address
// for photo upload onto a user's wall.
type WallUploadServer struct {
	Response struct {
		UploadURL string `json:"upload_url"`
		AlbumID int `json:"album_id"`
		UserID int `json:"user_id"`
	} `json:"response"`
}

// SavedWallPhoto describes info about
// saved photo on wall after being uploaded.
type SavedWallPhoto struct {
	Response []struct {
		Pid int `json:"pid"`
		ID string `json:"id"`
		Aid int `json:"aid"`
		OwnerID int `json:"owner_id"`
		Src string `json:"src"`
		SrcBig string `json:"src_big"`
		SrcSmall string `json:"src_small"`
		SrcXbig string `json:"src_xbig"`
		Width int `json:"width"`
		Height int `json:"height"`
		Text string `json:"text"`
		Created int `json:"created"`
	} `json:"response"`
}