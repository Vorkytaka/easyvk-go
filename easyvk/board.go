package easyvk

import (
	"fmt"
	"strconv"
)

// A Board describes a set of methods to work with topics.
// https://vk.com/dev/board
type Board struct {
	vk *VK
}

// BoardAddTopicParams provides fields for AddTopic params.
// https://vk.com/dev/board.addTopic
type BoardAddTopicParams struct {
	GroupID     uint
	Title       string
	Text        string
	FromGroup   bool
	Attachments string
}

// AddTopic creates a new topic on a community's discussion board.
// https://vk.com/dev/board.addTopic
func (b *Board) AddTopic(p BoardAddTopicParams) (int, error) {
	params := map[string]string{
		"group_id":    fmt.Sprint(p.GroupID),
		"title":       p.Title,
		"text":        p.Text,
		"from_group":  boolConverter(p.FromGroup),
		"attachments": p.Attachments,
	}
	resp, err := b.vk.Request("board.addTopic", params)
	if err != nil {
		return 0, err
	}
	topicID, err := strconv.Atoi(string(resp))
	if err != nil {
		return 0, err
	}
	return int(topicID), nil
}
