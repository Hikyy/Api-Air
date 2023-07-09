package requests

type UserLoginRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			Email       string `json:"email_address" validate:"email|min=6|required"`
			Password    string `json:"password" validate:"max=255|required"`
		} `json:"attributes"`
	} `json:"data"`
}
