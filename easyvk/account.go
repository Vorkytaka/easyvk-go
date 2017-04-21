package easyvk

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	notifyBit = 1 << iota
	friendsBit
	photosBit
	audioBit
	videoBit
	_
	_
	pagesBit
	_
	_
	statusBit
	notesBit
	messagesBit
	wallBit
	_
	adsBit
	offlineBit
	docsBit
	groupsBit
	notificationsBit
	statsBit
	_
	emailBit
	_
	_
	_
	_
	_
	marketBit
)

// An Account describes a set of methods
// to work with account.
type Account struct {
	vk *VK
}

// GetInfo returns current account info.
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

// GetProfileInfo returns the current account info.
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

// GetCounters returns values of user counters.
func (a *Account) GetCounters(filter string) (Counters, error) {
	params := map[string]string{"filter": filter }
	resp, err := a.vk.Request("account.getCounters", params)
	if err != nil {
		return Counters{}, err
	}
	var counters Counters
	err = json.Unmarshal(resp, &counters)
	if err != nil {
		return Counters{}, err
	}
	return counters, nil
}

// GetAppPermissions returns settings of the user in this application.
func (a *Account) GetAppPermissions(userID uint) (Permissions, error) {
	params := map[string]string{"user_id": fmt.Sprint(userID) }
	resp, err := a.vk.Request("account.getAppPermissions", params)
	if err != nil {
		return Permissions{}, err
	}

	permBits, err := strconv.ParseUint(string(resp), 10, 64)
	if err != nil {
		return Permissions{}, err
	}

	perm := Permissions{}

	perm.Notify = permBits&notifyBit != 0
	perm.Friends = permBits&friendsBit != 0
	perm.Photos = permBits&photosBit != 0
	perm.Audio = permBits&audioBit != 0
	perm.Video = permBits&videoBit != 0
	perm.Pages = permBits&pagesBit != 0
	perm.Status = permBits&statusBit != 0
	perm.Notes = permBits&notesBit != 0
	perm.Messages = permBits&messagesBit != 0
	perm.Wall = permBits&wallBit != 0
	perm.Ads = permBits&adsBit != 0
	perm.Offline = permBits&offlineBit != 0
	perm.Docs = permBits&docsBit != 0
	perm.Groups = permBits&groupsBit != 0
	perm.Notifications = permBits&notificationsBit != 0
	perm.Stats = permBits&statsBit != 0
	perm.Email = permBits&emailBit != 0
	perm.Market = permBits&marketBit != 0

	return perm, nil
}

// GetBanned returns a user's blacklist.
func (a *Account) GetBanned(offset, count uint) (Blacklist, error) {
	params := map[string]string{
		"offset": fmt.Sprint(offset),
		"count": fmt.Sprint(count),
	}
	resp, err := a.vk.Request("account.getBanned", params)
	if err != nil {
		return Blacklist{}, err
	}
	var list Blacklist
	err = json.Unmarshal(resp, &list)
	if err != nil {
		return Blacklist{}, err
	}
	return list, nil
}