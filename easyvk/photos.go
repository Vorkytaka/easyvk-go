package easyvk

import (
	"encoding/json"
	"fmt"
)

// A Photos describes a set of methods
// to work with photos.
type Photos struct {
	vk *VK
}

// GetWallUploadServer returns the server address for photo upload onto a user's wall.
func (p *Photos) GetWallUploadServer(groupID uint) (WallUploadServer, error) {
	params := map[string]string{"group_id": fmt.Sprint(groupID) }
	resp, err := p.vk.Request("photos.getWallUploadServer", params)
	if err != nil {
		return WallUploadServer{}, err
	}
	var server WallUploadServer
	err = json.Unmarshal(resp, &server)
	if err != nil {
		return WallUploadServer{}, err
	}
	return server, nil
}
