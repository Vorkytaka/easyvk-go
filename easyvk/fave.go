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
