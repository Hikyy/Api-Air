package requests

type SensorRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			RoomID  int `json:"room_id"`
			CmdType int `json:"cmd_type"`
		} `json:"attributes"`
	} `json:"data"`
}
