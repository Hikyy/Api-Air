package requests

//Todo: fix if value define in rule is int and if data is empty golang replace this by 0

import (
	"App/internal/common"
	"App/internal/helpers"
	"reflect"
	"strings"
)

type ruleDelimiterValue struct {
	first  string `operator:":"`
	second string `operator:"="`
}

type ruleType struct {
	email         string `operator:"" func:"Email"`
	integer       string `operator:"" func:"Integer"`
	string        string `operator:"" func:"String"`
	min           string `operator:"=" func:"Min"`
	max           string `operator:"=" func:"Max"`
	unique        string `operator:":" func:"Unique"`
	in            string `operator:":" func:"In"`
	nullable      string `operator:"" func:"Nullable"`
	between       string `operator:":" func:"Between"`
	date          string `operator:"" func:"Date"`
	required      string `operator:"" func:"Required"`
	required_with string `operator:":" func:"RequiredWith"`
}

func Rule(rule string, value reflect.Value) {
	for _, v := range strings.Split(rule, "|") {
		structValue := reflect.ValueOf(ruleType{})
		structType := structValue.Type()
		rule, valueRule := getKeyRule(v)
		fieldType, exists := structType.FieldByName(rule)
		if exists {
			data := common.Parameter{
				Data:      value,
				ValueRule: valueRule,
			}
			callFunction(data, fieldType.Tag.Get("func"))
		}
	}
}

func ToArray(p common.Parameter) []reflect.Value {
	return []reflect.Value{reflect.ValueOf(p)}
}

func callFunction(fn common.Parameter, name string) {
	method := reflect.ValueOf(&helpers.FuncCall{}).MethodByName(name)
	if method.IsValid() {
		method.Call(ToArray(fn))
	}
}

func getKeyRule(value string) (string, string) {
	valueOf := reflect.ValueOf(ruleDelimiterValue{})
	for i := 0; i < valueOf.NumField(); i++ {
		field := valueOf.Type().Field(i)
		if field.Tag.Get("operator") != "" && strings.Contains(value, field.Tag.Get("operator")) {
			split := strings.Split(value, field.Tag.Get("operator"))
			return split[0], split[1]
		}
	}
	return value, value
}