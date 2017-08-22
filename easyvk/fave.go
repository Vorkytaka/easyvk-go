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

// A FaveGetUsersResponse describes a list of users
// whom the current user has bookmarked.
// https://vk.com/dev/fave.getUsers
type FaveGetUsersResponse struct {
	Count int `json:"count"`
	Items []UserObject
}

// GetUsers returns a list of users whom the current user has bookmarked.
// https://vk.com/dev/fave.getUsers
func (f *Fave) GetUsers(offset, count uint) (FaveGetUsersResponse, error) {
	params := map[string]string{
		"offset": fmt.Sprint(offset),
		"count":  fmt.Sprint(count),
	}
	resp, err := f.vk.Request("fave.getUsers", params)
	if err != nil {
		return FaveGetUsersResponse{}, err
	}
	var users FaveGetUsersResponse
	err = json.Unmarshal(resp, &users)
	if err != nil {
		return FaveGetUsersResponse{}, err
	}

	return users, nil
}

// A FaveGetLinksResponse describes a list of links
// that the current user has bookmarked.
// https://vk.com/dev/fave.getLinks
type FaveGetLinksResponse struct {
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
func (f *Fave) GetLinks(offset, count uint) (FaveGetLinksResponse, error) {
	params := map[string]string{
		"offset": fmt.Sprint(offset),
		"count":  fmt.Sprint(count),
	}
	resp, err := f.vk.Request("fave.getLinks", params)
	if err != nil {
		return FaveGetLinksResponse{}, err
	}
	var links FaveGetLinksResponse
	err = json.Unmarshal(resp, &links)
	if err != nil {
		return FaveGetLinksResponse{}, err
	}

	return links, nil
}

// A FaveGetPhotosResponse describes a list of photos
// that the current user has bookmarked.
// https://vk.com/dev/fave.getPhotos
type FaveGetPhotosResponse struct {
	Count int `json:"count"`
	Items []PhotoObject
}

// GetPhotos returns a list of photos that the current user has bookmarked.
// https://vk.com/dev/fave.getPhotos
func (f *Fave) GetPhotos(offset, count uint) (FaveGetPhotosResponse, error) {
	params := map[string]string{
		"offset":      fmt.Sprint(offset),
		"count":       fmt.Sprint(count),
		"photo_sizes": "1",
	}
	resp, err := f.vk.Request("fave.getPhotos", params)
	if err != nil {
		return FaveGetPhotosResponse{}, err
	}
	var links FaveGetPhotosResponse
	err = json.Unmarshal(resp, &links)
	if err != nil {
		return FaveGetPhotosResponse{}, err
	}

	return links, nil
}

// A FaveGetVideosResponse describes a list of videos
// that the current user has bookmarked.
// https://vk.com/dev/fave.getVideos
type FaveGetVideosResponse struct {
	Count int `json:"count"`
	Items []VideoObject
}

// GetVideos returns a list of videos that the current user has bookmarked.
// https://vk.com/dev/fave.getVideos
func (f *Fave) GetVideos(offset, count uint) (FaveGetVideosResponse, error) {
	params := map[string]string{
		"offset":   fmt.Sprint(offset),
		"count":    fmt.Sprint(count),
		"extended": "1",
	}
	resp, err := f.vk.Request("fave.getVideos", params)
	if err != nil {
		return FaveGetVideosResponse{}, err
	}
	var links FaveGetVideosResponse
	err = json.Unmarshal(resp, &links)
	if err != nil {
		return FaveGetVideosResponse{}, err
	}

	return links, nil
}
