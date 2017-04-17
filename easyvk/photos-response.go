package easyvk

type WallUploadServer struct {
	Response struct {
		UploadURL string `json:"upload_url"`
		AlbumID   int `json:"aid"`
		UserID    int `json:"mid"`
	} `json:"response"`
}
