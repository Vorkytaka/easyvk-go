package easyvk

import (
	"encoding/json"
	"fmt"
)

// A Photos describes a set of methods
// to work with photos.
// https://vk.com/dev/photos
type Photos struct {
	vk *VK
}

// PhotosGetWallUploadServerResponse describes the server address
// for photo upload onto a user's wall.
// https://vk.com/dev/photos.getWallUploadServer
type PhotosGetWallUploadServerResponse struct {
	UploadURL string `json:"upload_url"`
	AlbumID   int `json:"album_id"`
	UserID    int `json:"user_id"`
}

// GetWallUploadServer returns the server address for photo upload onto a user's wall.
// https://vk.com/dev/photos.getWallUploadServer
func (p *Photos) GetWallUploadServer(groupID uint) (PhotosGetWallUploadServerResponse, error) {
	params := map[string]string{"group_id": fmt.Sprint(groupID) }
	resp, err := p.vk.Request("photos.getWallUploadServer", params)
	if err != nil {
		return PhotosGetWallUploadServerResponse{}, err
	}
	var server PhotosGetWallUploadServerResponse
	err = json.Unmarshal(resp, &server)
	if err != nil {
		return PhotosGetWallUploadServerResponse{}, err
	}
	return server, nil
}

// PhotosSaveWallPhotoParams provides structure for
// parameters for saveWallPhoto method.
// https://vk.com/dev/photos.saveWallPhoto
type PhotosSaveWallPhotoParams struct {
	UserID  uint
	GroupID uint
	Photo   string
	Hash    string
	Caption string
	Server  int
	Lat     float64
	Long    float64
}

// SaveWallPhoto saves a photo to a user's or community's wall after being uploaded.
// For upload look at file upload.go.
// https://vk.com/dev/photos.saveWallPhoto
func (p *Photos) SaveWallPhoto(par PhotosSaveWallPhotoParams) ([]PhotoObject, error) {
	params := map[string]string{
		"user_id":   fmt.Sprint(par.UserID),
		"group_id":  fmt.Sprint(par.GroupID),
		"photo":     par.Photo,
		"hash":      par.Hash,
		"caption":   par.Caption,
		"server":    fmt.Sprint(par.Server),
		"latitude":  fmt.Sprint(par.Lat),
		"longitude": fmt.Sprint(par.Long),
	}
	resp, err := p.vk.Request("photos.saveWallPhoto", params)
	if err != nil {
		return nil, err
	}

	var info []PhotoObject
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
