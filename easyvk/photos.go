package easyvk

import (
	"encoding/json"
	"fmt"
)

type Photos struct {
	vk *VK
}

func (p *Photos) GetWallUploadServer(group_id uint) (WallUploadServer, error) {
	params := map[string]string{"group_id": fmt.Sprint(group_id) }
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
