package easyvk

import (
	"fmt"
	"encoding/json"
)

// A Likes describes a set of methods to work with likes.
// https://vk.com/dev/likes
type Likes struct {
	vk *VK
}

type likeType string

const (
	// PostLikeType - post on user or community wall
	PostLikeType likeType = "post"
	// CommentLikeType - comment on a wall post
	CommentLikeType likeType = "comment"
	// PhotoLikeType - photo
	PhotoLikeType likeType = "photo"
	// AudioLikeType - audio
	AudioLikeType likeType = "audio"
	// VideoLikeType - video
	VideoLikeType likeType = "video"
	// NoteLikeType - note
	NoteLikeType likeType = "note"
	// MarketLikeType - market
	MarketLikeType likeType = "market"
	// PhotoCommentLikeType - comment on the photo
	PhotoCommentLikeType likeType = "photo_comment"
	// VideoCommentLikeType - comment on the video
	VideoCommentLikeType likeType = "video_comment"
	// TopicCommentLikeType - comment in the discussion
	TopicCommentLikeType likeType = "topic_comment"
	// MarketCommentLikeType - comment on the market
	MarketCommentLikeType likeType = "market_comment"
)

// Add adds the specified object to the Likes list of the current user.
// https://vk.com/dev/likes.add
func (l *Likes) Add(t likeType, ownerID int, itemID uint, accessKey string) (int, error) {
	params := map[string]string{
		"type":       string(t),
		"owner_id":   fmt.Sprint(ownerID),
		"item_id":    fmt.Sprint(itemID),
		"access_key": accessKey,
	}
	resp, err := l.vk.Request("likes.add", params)
	if err != nil {
		return 0, err
	}
	var likes struct {
		Likes int `json:"likes"`
	}
	err = json.Unmarshal(resp, &likes)
	if err != nil {
		return 0, err
	}
	return likes.Likes, nil
}
