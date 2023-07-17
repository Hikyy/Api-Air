package resources

type RoomResource struct {
	Data struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			RoomNumber int `json:"room_number"`
			FloorId    int `json:"floor_id"`
		} `json:"attributes"`
	} `json:"data"`
}
