package easyvk

import (
	"encoding/json"
	"fmt"
)

// A Photos describes a set of methods
// to work with photos.
type Photos struct {
	vk *VK
}

// WallUploadServer describes the server address
// for photo upload onto a user's wall.
type WallUploadServer struct {
	UploadURL string `json:"upload_url"`
	AlbumID   int `json:"album_id"`
	UserID    int `json:"user_id"`
}

// GetWallUploadServer returns the server address for photo upload onto a user's wall.
func (p *Photos) GetWallUploadServer(groupID uint) (WallUploadServer, error) {
	params := map[string]string{"group_id": fmt.Sprint(groupID) }
	resp, err := p.vk.Request("photos.getWallUploadServer", params)
	if err != nil {
		return WallUploadServer{}, err
	}
	var server WallUploadServer
	err = json.Unmarshal(resp, &server)
	if err != nil {
		return WallUploadServer{}, err
	}
	return server, nil
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

// SaveWallPhoto saves a photo to a user's or community's wall after being uploaded.
// For upload look at file upload.go.
func (p *Photos) SaveWallPhoto(userID, groupID uint, photo, hash, caption string, server int, latitude, longitude float64) (SavedWallPhoto, error) {
	params := map[string]string{
		"user_id":   fmt.Sprint(userID),
		"group_id":  fmt.Sprint(groupID),
		"photo":     photo,
		"hash":      hash,
		"caption":   caption,
		"server":    fmt.Sprint(server),
		"latitude":  fmt.Sprint(latitude),
		"longitude": fmt.Sprint(longitude),
	}
	resp, err := p.vk.Request("photos.saveWallPhoto", params)
	if err != nil {
		return SavedWallPhoto{}, err
	}

	var info SavedWallPhoto
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return SavedWallPhoto{}, err
	}
	return info, nil
}
