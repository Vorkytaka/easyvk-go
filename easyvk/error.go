package easyvk

// An Error describes vk errors info.
type Error struct {
	Code    int `json:"error_code"`
	Message string `json:"error_msg"`
	RequestParams []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params"`
}