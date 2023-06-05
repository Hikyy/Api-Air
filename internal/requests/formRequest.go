package requests

import (
	"fmt"
	"net/http"
)

type Rule struct {
	Result  map[string]interface{}
	Request http.ResponseWriter
	Struct  entity
}

type entity struct {
}

func (f Rule) Validated() {
	for k, v := range f.Result {
		fmt.Println(k, v)
	}
}
