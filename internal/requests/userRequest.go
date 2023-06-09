package requests

type StoreUserRequest struct {
	Data struct {
		Type       string `json:"type"`
		Attributes struct {
			EmailAddress string `json:"email_address" validate:"email|max=255|required"`
			FullName     string `json:"full_name" validate:"max=255|required"`
			Password     string `json:"password" validate:"max=255|required"`
			RedirectURL  string `json:"redirect_url" validate:"max=255|required"`
		} `json:"attributes"`
	} `json:"data"`
}
