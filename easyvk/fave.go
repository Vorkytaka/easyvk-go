package easyvk

import (
	"fmt"
	"encoding/json"
)

// A Fave describes a set of methods
// to work with faves.
// https://vk.com/dev/fave
type Fave struct {
	vk *VK
}

// A FaveUsersResponse describes a list of users
// whom the current user has bookmarked.
// https://vk.com/dev/fave.getUsers
type FaveUsersResponse struct {
	Count int `json:"count"`
	Items []UserObject
}

// GetUsers returns a list of users whom the current user has bookmarked.
// https://vk.com/dev/fave.getUsers
func (f *Fave) GetUsers(offset, count uint) (FaveUsersResponse, error) {
	params := map[string]string{
		"offset": fmt.Sprint(offset),
		"count":  fmt.Sprint(count),
	}
	resp, err := f.vk.Request("fave.getUsers", params)
	if err != nil {
		return FaveUsersResponse{}, err
	}
	var users FaveUsersResponse
	err = json.Unmarshal(resp, &users)
	if err != nil {
		return FaveUsersResponse{}, err
	}

	return users, nil
}

// A FaveLinksResponse describes a list of links
// that the current user has bookmarked.
// https://vk.com/dev/fave.getLinks
type FaveLinksResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID          string `json:"id"`
		URL         string `json:"url"`
		Title       string `json:"title"`
		Description string `json:"description"`
		Photo50     string `json:"photo_50"`
		Photo100    string `json:"photo_100"`
		Photo200    string `json:"photo_200"`
	} `json:"items"`
}

// GetLinks returns a list of links that the current user has bookmarked.
// https://vk.com/dev/fave.getLinks
func (f *Fave) GetLinks(offset, count uint) (FaveLinksResponse, error) {
	params := map[string]string{
		"offset": fmt.Sprint(offset),
		"count":  fmt.Sprint(count),
	}
	resp, err := f.vk.Request("fave.getLinks", params)
	if err != nil {
		return FaveLinksResponse{}, err
	}
	var links FaveLinksResponse
	err = json.Unmarshal(resp, &links)
	if err != nil {
		return FaveLinksResponse{}, err
	}

	return links, nil
}

// A FavePhotosResponse describes a list of photos
// that the current user has bookmarked.
// https://vk.com/dev/fave.getPhotos
type FavePhotosResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID        int `json:"id"`
		AlbumID   int `json:"album_id"`
		OwnerID   int `json:"owner_id"`
		UserID    int `json:"user_id"`
		Photo75   string `json:"photo_75"`
		Photo130  string `json:"photo_130"`
		Photo604  string `json:"photo_604"`
		Photo807  string `json:"photo_807"`
		Photo1280 string `json:"photo_1280"`
		Photo2560 string `json:"photo_2560"`
		Width     int `json:"width"`
		Height    int `json:"height"`
		Text      string `json:"text"`
		Date      int `json:"date"`
		PostID    int `json:"post_id"`
		AccessKey string `json:"access_key"`
	} `json:"items"`
}

// GetPhotos returns a list of photos that the current user has bookmarked.
// https://vk.com/dev/fave.getPhotos
func (f *Fave) GetPhotos(offset, count uint) (FavePhotosResponse, error) {
	params := map[string]string{
		"offset":      fmt.Sprint(offset),
		"count":       fmt.Sprint(count),
		"photo_sizes": "0",
	}
	resp, err := f.vk.Request("fave.getPhotos", params)
	if err != nil {
		return FavePhotosResponse{}, err
	}
	var links FavePhotosResponse
	err = json.Unmarshal(resp, &links)
	if err != nil {
		return FavePhotosResponse{}, err
	}

	return links, nil
}

// A FaveVideosResponse describes a list of videos
// that the current user has bookmarked.
// https://vk.com/dev/fave.getVideos
type FaveVideosResponse struct {
	Count int `json:"count"`
	Items []struct {
		ID          int `json:"id"`
		OwnerID     int `json:"owner_id"`
		Title       string `json:"title"`
		Duration    int `json:"duration"`
		Description string `json:"description"`
		Date        int `json:"date"`
		Comments    int `json:"comments"`
		Views       int `json:"views"`
		Width       int `json:"width"`
		Height      int `json:"height"`
		Photo130    string `json:"photo_130"`
		Photo320    string `json:"photo_320"`
		Photo800    string `json:"photo_800"`
		CanAdd      int `json:"can_add"`
	} `json:"items"`
}

// GetVideos returns a list of videos that the current user has bookmarked.
// https://vk.com/dev/fave.getVideos
func (f *Fave) GetVideos(offset, count uint) (FaveVideosResponse, error) {
	params := map[string]string{
		"offset":   fmt.Sprint(offset),
		"count":    fmt.Sprint(count),
		"extended": "0",
	}
	resp, err := f.vk.Request("fave.getVideos", params)
	if err != nil {
		return FaveVideosResponse{}, err
	}
	var links FaveVideosResponse
	err = json.Unmarshal(resp, &links)
	if err != nil {
		return FaveVideosResponse{}, err
	}

	return links, nil
}
