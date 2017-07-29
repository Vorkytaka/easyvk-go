package easyvk

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// A Status describes a set of methods
// to work with status.
type Status struct {
	vk *VK
}

// Get returns data required to show the
// status of a user or community
func (s *Status) Get(id int) (string, error) {
	params := map[string]string{"user_id": fmt.Sprint(id) }
	resp, err := s.vk.Request("status.get", params)
	if err != nil {
		return "", err
	}
	var status struct {
		Text string `json:"text"`
	}
	err = json.Unmarshal(resp, &status)
	if err != nil {
		return "", err
	}
	return status.Text, nil
}

// Set a new status for the current user
func (s *Status) Set(text string, id int) (bool, error) {
	params := map[string]string{
		"text":     text,
		"group_id": fmt.Sprint(id),
	}
	resp, err := s.vk.Request("status.set", params)
	if err != nil {
		return false, err
	}
	ok, err := strconv.ParseUint(string(resp), 10, 8)
	if err != nil {
		return false, err
	}
	return ok == 1, nil
}
