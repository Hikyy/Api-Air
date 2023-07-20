package resources


type ActuatorResource struct {
    Data struct {
        Type       string `json:"type"`
        Id              int    `json:"id"`
        Attributes struct {
            ActuatorName    string `json:"actuator_name"`
            ActuatorCommand string `json:"actuator_command"`
            DataKey         string `json:"data_key"`
            RoomId          int    `json:"room_id"`
        } `json:"attributes"`
    } `json:"data"`
}