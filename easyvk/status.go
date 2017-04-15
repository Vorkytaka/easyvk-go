package easyvk

import (
	"encoding/json"
	"fmt"
)

type Status struct {
	vk *VK
}

func (s *Status) Get(id int) (string, error) {
	params := map[string]string{"user_id": fmt.Sprint(id) }
	resp, err := s.vk.Request("status.get", params)
	if err != nil {
		return "", err
	}
	var status StatusInfo
	err = json.Unmarshal(resp, &status)
	if err != nil {
		return "", err
	}
	return status.Response.Text, nil
}

func (s *Status) Set(text string, id int) (bool, error) {
	params := map[string]string{
		"text": text,
		"group_id": fmt.Sprint(id),
	}
	resp, err := s.vk.Request("status.set", params)
	if err != nil {
		return false, err
	}
	var status Response
	err = json.Unmarshal(resp, &status)
	if err != nil {
		return false, err
	}
	return status.Response == 1, nil
}
