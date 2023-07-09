package requests

type UpdateUserRequest struct {
	Data struct {
		Type       string `json:"type"`
		Id         string `json:"id"`
		Attributes struct {
			Email      string `json:"email_address" validate:"email|min=6"`
			Firstname  string `json:"firstname" validate:"max=255"`
			Lastname   string `json:"lastname" validate:"max=255"`
			Password   string `json:"password" validate:"max=255"`
			Group_name string `json:"group_name"`
		} `json:"attributes"`
	} `json:"data"`
}
