package easyvk

import (
	"encoding/json"
)

type Account struct {
	vk *VK
}

func (a *Account) GetInfo(fields string) Info {
	params := map[string]string{"fields": fields }
	resp, err := a.vk.Request("account.getInfo", params)
	if err != nil {

	}
	var info Info
	err = json.Unmarshal(resp, &info)
	if err != nil {

	}
	return info
}

func (a *Account) GetProfileInfo() ProfileInfo {
	resp, err := a.vk.Request("account.getProfileInfo", nil)
	if err != nil {

	}
	var info ProfileInfo
	err = json.Unmarshal(resp, &info)
	if err != nil {

	}
	return info
}
