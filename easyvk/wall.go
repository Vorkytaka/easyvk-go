package easyvk

import (
	"fmt"
	"encoding/json"
)

// A Wall describes a set of methods
// to work with wall.
// https://vk.com/dev/wall
type Wall struct {
	vk *VK
}

// Post adds a new post on a user wall or community wall.
// Can also be used to publish suggested or scheduled posts.
// Returns id of created post.
// https://vk.com/dev/wall.post
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
	params["friends_only"] = boolConverter(friendsOnly)
	params["from_group"] = boolConverter(fromGroup)
	params["signed"] = boolConverter(signed)
	params["mark_as_ads"] = boolConverter(markAsAds)
	params["ads_promoted_stealth"] = boolConverter(adsPromotedStealth)

	resp, err := w.vk.Request("wall.post", params)
	if err != nil {
		return 0, err
	}
	var info struct {
		PostID int `json:"post_id"`
	}

	err = json.Unmarshal(resp, &info)
	if err != nil {
		return 0, err
	}
	return info.PostID, nil
}
