package requests

type SensorRequest struct {
	CmdId              int    `json:"cmd_id"`
	DestinationAddress string `json:"destination_address"`
	AckFlags           int    `json:"ack_flags"`
	CmdType            int    `json:"cmd_type"`
}
