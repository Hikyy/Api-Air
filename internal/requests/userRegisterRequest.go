package requests

type StoreUserRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			Email      string `json:"email_address" validate:"email|min=6|required"`
			Firstname  string `json:"firstname" validate:"max=255|required"`
			Lastname   string `json:"lastname" validate:"max=255|required"`
			Password   string `json:"password" validate:"max=255|required"`
			Group_name string `json:"default:'user'"`
		} `json:"attributes"`
	} `json:"data"`
}
