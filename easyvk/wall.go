package easyvk

import (
	"fmt"
	"encoding/json"
)

// A Wall describes a set of methods
// to work with wall.
type Wall struct {
	vk *VK
}

// Post adds a new post on a user wall or community wall.
// Can also be used to publish suggested or scheduled posts.
// Returns id of post.
func (w *Wall) Post(ownerID int,
	friendsOnly, fromGroup, signed, markAsAds, adsPromotedStealth bool,
	message, attachments, services, guid string,
	publishDate, placeID, postID uint,
	lat, long float64) (int, error) {

	params := map[string]string{
		"owner_id":     fmt.Sprint(ownerID),
		"message":      message,
		"attachments":  attachments,
		"services":     services,
		"guid":         guid,
		"publish_date": fmt.Sprint(publishDate),
		"place_id":     fmt.Sprint(placeID),
		"post_id":      fmt.Sprint(postID),
		"lat":          fmt.Sprint(lat),
		"long":         fmt.Sprint(long),
	}
	if friendsOnly {
		params["friends_only"] = "1"
	} else {
		params["friends_only"] = "0"
	}
	if fromGroup {
		params["from_group"] = "1"
	} else {
		params["from_group"] = "0"
	}
	if signed {
		params["signed"] = "1"
	} else {
		params["signed"] = "0"
	}
	if markAsAds {
		params["mark_as_ads"] = "1"
	} else {
		params["mark_as_ads"] = "0"
	}
	if adsPromotedStealth {
		params["ads_promoted_stealth"] = "1"
	} else {
		params["ads_promoted_stealth"] = "0"
	}

	resp, err := w.vk.Request("wall.post", params)
	if err != nil {
		return 0, nil
	}
	var info wallPost
	err = json.Unmarshal(resp, &info)
	if err != nil {
		return 0, nil
	}
	return info.Response.PostID, nil
}
