package easyvk

// StatusInfo describes status text.
type StatusInfo struct {
	Response struct {
		Text string `json:"text"`
	} `json:"response"`
}
