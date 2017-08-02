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

// WallPostParams provides structure for post method.
// https://vk.com/dev/wall.post
type WallPostParams struct {
	OwnerID            int
	FriendsOnly        bool
	FromGroup          bool
	Signed             bool
	MarkAsAds          bool
	AdsPromotedStealth bool
	Message            string
	Attachments        string
	Services           string
	GUID               string
	PublishDate        uint
	PlaceID            uint
	PostID             uint
	Lat                float64
	Long               float64
}

// Post adds a new post on a user wall or community wall.
// Can also be used to publish suggested or scheduled posts.
// Returns id of created post.
// https://vk.com/dev/wall.post
func (w *Wall) Post(p WallPostParams) (int, error) {

	params := map[string]string{
		"owner_id":             fmt.Sprint(p.OwnerID),
		"message":              p.Message,
		"attachments":          p.Attachments,
		"services":             p.Services,
		"guid":                 p.GUID,
		"publish_date":         fmt.Sprint(p.PublishDate),
		"place_id":             fmt.Sprint(p.PlaceID),
		"post_id":              fmt.Sprint(p.PostID),
		"lat":                  fmt.Sprint(p.Lat),
		"long":                 fmt.Sprint(p.Long),
		"friends_only":         boolConverter(p.FriendsOnly),
		"from_group":           boolConverter(p.FromGroup),
		"signed":               boolConverter(p.Signed),
		"mark_as_ads":          boolConverter(p.MarkAsAds),
		"ads_promoted_stealth": boolConverter(p.AdsPromotedStealth),
	}

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
