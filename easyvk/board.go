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

// CloseTopic closes a topic on a community's discussion board so that comments cannot be posted.
// https://vk.com/dev/board.closeTopic
func (b *Board) CloseTopic(groupID, topicID uint) (bool, error) {
	params := map[string]string{
		"group_id": fmt.Sprint(groupID),
		"topic_id": fmt.Sprint(topicID),
	}
	resp, err := b.vk.Request("board.closeTopic", params)
	if err != nil {
		return false, err
	}
	ok, err := strconv.ParseUint(string(resp), 10, 8)
	if err != nil {
		return false, err
	}
	return ok == 1, nil
}
