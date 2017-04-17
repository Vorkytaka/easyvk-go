package easyvk

// WallUploadServer describes the server address
// for photo upload onto a user's wall.
type WallUploadServer struct {
	Response struct {
		UploadURL string `json:"upload_url"`
		AlbumID   int `json:"aid"`
		UserID    int `json:"mid"`
	} `json:"response"`
}
