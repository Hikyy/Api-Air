package resources

type UserResource struct {
	Data struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			Firstname  string `json:"firstname"`
			Lastname   string `json:"lastname"`
			Email      string `json:"email"`
			Group_name string `json:"group_name"`
		} `json:"attributes"`
	} `json:"data"`
}
