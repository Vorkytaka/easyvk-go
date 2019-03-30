package easyvk

import (
	"encoding/json"
	"fmt"
	"strings"
)

// A Friends describes a set of methods
// to work with friends.
// https://vk.com/dev/friends
type Friends struct {
	vk *VK
}

// A FriendsGetResponse describes a list of friends
// of the user.
// https://vk.com/dev/friends.get
type FriendsGetResponse struct {
	Count int `json:"count"`
	Items []UserObject
}

// Get returns a list of friends of the user.
//https://vk.com/dev/friends.get
func (f *Friends) Get(userId int, fields []string) (FriendsGetResponse, error) {
	var fieldsWithId = fields
	if len(fields) == 0 {
		fieldsWithId = append(fields, "id")
	}

	params := map[string]string{
		"user_id": fmt.Sprint(userId),
		"fields":  strings.Join(fieldsWithId,","),
	}
	resp, err := f.vk.Request("friends.get", params)
	if err != nil {
		return FriendsGetResponse{}, err
	}
	var users FriendsGetResponse
	err = json.Unmarshal(resp, &users)
	if err != nil {
		return FriendsGetResponse{}, err
	}

	return users, nil
}

// Get returns a list of friends of the user.
//https://vk.com/dev/friends.get
func (f *Friends) Get(userId int) (FriendsGetResponse, error) {
	return f.Get(userId, []string{})
}
