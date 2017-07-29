package easyvk

import (
	"encoding/json"
	"fmt"
	"strconv"
)

const (
	notifyBit        = 1 << iota
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

// An Info describes a set of user's info.
type Info struct {
	Country         string `json:"country"`
	HTTPS           int `json:"https_required"`
	TwoFactor       int `json:"2fa_required"`
	OwnPostsDefault int `json:"own_posts_default"`
	NoWallReplies   int `json:"no_wall_replies"`
	Intro           int `json:"intro"`
	Lang            int `json:"lang"`
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

// A ProfileInfo describes a set of user's info.
type ProfileInfo struct {
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	ScreenName      string `json:"screen_name"`
	Sex             int `json:"sex"`
	Relation        int `json:"relation"`
	Birthday        string `json:"bdate"`
	BirthVisibility int `json:"bdate_visibility"`
	Hometown        string `json:"home_town"`
	Status          string `json:"status"`
	Phone           string `json:"phone"`
	Country struct {
		ID    int `json:"id"`
		Title string `json:"title"`
	} `json:"country"`
	City struct {
		ID    int `json:"id"`
		Title string `json:"title"`
	} `json:"city"`
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

// A Counters describes a set of user's counters.
type Counters struct {
	Friends            int `json:"friends"`
	FriendsSuggestions int `json:"friends_suggestions"`
	Messages           int `json:"messages"`
	Photos             int `json:"photos"`
	Videos             int `json:"videos"`
	Gifts              int `json:"gifts"`
	Events             int `json:"events"`
	Groups             int `json:"groups"`
	Notifications      int `json:"notifications"`
	SDK                int `json:"sdk"`
	AppRequests        int `json:"app_requests"`
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

// A Permissions describes a set of app's permissions.
type Permissions struct {
	Notify        bool
	Friends       bool
	Photos        bool
	Audio         bool
	Video         bool
	Pages         bool
	Status        bool
	Notes         bool
	Messages      bool
	Wall          bool
	Ads           bool
	Offline       bool
	Docs          bool
	Groups        bool
	Notifications bool
	Stats         bool
	Email         bool
	Market        bool
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

// A Blacklist describes a user's blacklist.
type Blacklist struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID          int `json:"id"`
			FirstName   string `json:"first_name"`
			LastName    string `json:"last_name"`
			Deactivated string `json:"deactivated"`
		} `json:"items"`
	} `json:"response"`
}

// GetBanned returns a user's blacklist.
func (a *Account) GetBanned(offset, count uint) (Blacklist, error) {
	params := map[string]string{
		"offset": fmt.Sprint(offset),
		"count":  fmt.Sprint(count),
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

// BanUser adds user to the banlist.
func (a *Account) BanUser(userID uint) (bool, error) {
	params := map[string]string{
		"user_id": fmt.Sprint(userID),
	}
	resp, err := a.vk.Request("account.banUser", params)
	if err != nil {
		return false, err
	}
	ok, err := strconv.ParseUint(string(resp), 10, 8)
	if err != nil {
		return false, err
	}
	return ok == 1, nil
}

// UnbanUser deletes user from the blacklist.
func (a *Account) UnbanUser(userID uint) (bool, error) {
	params := map[string]string{
		"user_id": fmt.Sprint(userID),
	}
	resp, err := a.vk.Request("account.unbanUser", params)
	if err != nil {
		return false, err
	}
	ok, err := strconv.ParseUint(string(resp), 10, 8)
	if err != nil {
		return false, err
	}
	return ok == 1, nil
}
