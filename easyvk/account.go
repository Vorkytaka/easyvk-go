package easyvk

import (
	"encoding/json"
)

type Account struct {
	vk *VK
}

func (a *Account) GetInfo(fields string) (Info, error) {
	params := map[string]string{"fields": fields }
	resp, err := a.vk.Request("account.getInfo", params)
	if err != nil {
		return Info{}, err
	}
	var info Info
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return Info{}, err
	}
	return info, nil
}

func (a *Account) GetProfileInfo() (ProfileInfo, error) {
	resp, err := a.vk.Request("account.getProfileInfo", nil)
	if err != nil {
		return ProfileInfo{}, err
	}
	var info ProfileInfo
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return ProfileInfo{}, err
	}
	return info, nil
}
