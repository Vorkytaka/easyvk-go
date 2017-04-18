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
// Also this func return count of all bookmarked users.
func (f *Fave) GetUsers(offset, count uint) (FaveUsers, uint, error) {
	params := map[string]string{
		"offset": fmt.Sprint(offset),
		"count":  fmt.Sprint(count),
	}
	resp, err := f.vk.Request("fave.getUsers", params)
	if err != nil {
		return FaveUsers{}, 0, err
	}
	nextResp, quantity, err := f.vk.ResponseWithCount(resp)
	var users FaveUsers
	err = json.Unmarshal(nextResp, &users)
	if err != nil {
		return FaveUsers{}, 0, err
	}

	return users, quantity, nil
}

// GetLinks returns a list of links that the current user has bookmarked.
// Also this func return count of all bookmarked links.
func (f *Fave) GetLinks(offset, count uint) (FaveLinks, uint, error) {
	params := map[string]string{
		"offset": fmt.Sprint(offset),
		"count":  fmt.Sprint(count),
	}
	resp, err := f.vk.Request("fave.getLinks", params)
	if err != nil {
		return FaveLinks{}, 0, err
	}
	nextResp, quantity, err := f.vk.ResponseWithCount(resp)
	var links FaveLinks
	err = json.Unmarshal(nextResp, &links)
	if err != nil {
		return FaveLinks{}, 0, err
	}

	return links, quantity, nil
}