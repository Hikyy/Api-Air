package requests

type CreateResourceRequest struct {
	EmailAddress string `json:"email_address"`
	Password     string `json:"password"`
}

func StoreUserRequest() map[string]interface{} {
	//packageType := reflect.TypeOf(fmt.Printf).PkgPath()
	//pkg, _ := reflect.Import(StoreUserRequest)
	//fmt.Println(pkg)
	return map[string]interface{}{
		"data": map[string]interface{}{
			"type": "in:users|required|string",
			"attributes": map[string]interface{}{
				"email_address": "min:8|max:255|required|string",
				"password":      "min:8|max:255|required|string",
			},
		},
	}
}
