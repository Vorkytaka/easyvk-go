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
