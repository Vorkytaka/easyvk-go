package easyvk

// A PhotoWall describes an info
// about uploaded photo.
type PhotoWall struct {
	Server int `json:"server"`
	Photo string `json:"photo"`
	Hash string `json:"hash"`
}