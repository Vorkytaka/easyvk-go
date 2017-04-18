package easyvk

import (
	"fmt"
	"encoding/json"
)

// A Fave describes a set of methods
// to work with faves.
type Fave struct {
	vk *VK
}

// GetUsers returns a list of users whom the current user has bookmarked.
func (f *Fave) GetUsers(offset, count uint) (FaveUsers, error) {
	params := map[string]string{
		"offset": fmt.Sprint(offset),
		"count":  fmt.Sprint(count),
	}
	resp, err := f.vk.Request("fave.getUsers", params)
	if err != nil {
		return FaveUsers{}, err
	}
	nextResp, count, err := f.vk.ResponseWithCount(resp)
	var users FaveUsers
	err = json.Unmarshal(nextResp, &users)
	if err != nil {
		return FaveUsers{}, err
	}

	return users, nil
}

// GetLinks returns a list of links that the current user has bookmarked.
func (f *Fave) GetLinks(offset, count uint) (FaveLinks, error) {
	params := map[string]string{
		"offset": fmt.Sprint(offset),
		"count":  fmt.Sprint(count),
	}
	resp, err := f.vk.Request("fave.getLinks", params)
	if err != nil {
		return FaveLinks{}, err
	}
	nextResp, count, err := f.vk.ResponseWithCount(resp)
	var links FaveLinks
	err = json.Unmarshal(nextResp, &links)
	if err != nil {
		return FaveLinks{}, err
	}

	return links, nil
}
