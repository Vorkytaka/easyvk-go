package easyvk

// WallUploadServer describes the server address
// for photo upload onto a user's wall.
type WallUploadServer struct {
	UploadURL string `json:"upload_url"`
	AlbumID   int `json:"album_id"`
	UserID    int `json:"user_id"`
}

// SavedWallPhoto describes info about
// saved photo on wall after being uploaded.
type SavedWallPhoto []struct {
	ID       int `json:"id"`
	AlbumID  int `json:"album_id"`
	OwnerID  int `json:"owner_id"`
	Photo75  string `json:"photo_75"`
	Photo130 string `json:"photo_130"`
	Photo604 string `json:"photo_604"`
	Photo807 string `json:"photo_807"`
	Width    int `json:"width"`
	Height   int `json:"height"`
	Text     string `json:"text"`
	Date     int `json:"date"`
}
