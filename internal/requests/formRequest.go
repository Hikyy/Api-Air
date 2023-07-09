package requests

//Todo: fix if value define in rule is int and if data is empty golang replace this by 0

import (
	"App/internal/common"
	"App/internal/helpers"
	"fmt"
	"reflect"
	"strings"
)

type ruleDelimiterValue struct {
	first  string `operator:":"`
	second string `operator:"="`
}

type RuleType struct {
	email         string `operator:"" func:"Email" message:"n'est pas un email valide"`
	integer       string `operator:"" func:"Integer" message:"n'est pas un entier"`
	string        string `operator:"" func:"String" message:"n'est pas une chaine de caractère"`
	min           string `operator:"=" func:"Min" message:"doit être supérieur à %v"`
	max           string `operator:"=" func:"Max" message:"doit être inférieur à %v"`
	unique        string `operator:":" func:"Unique" message:"existe déjà"`
	in            string `operator:":" func:"In" message:"n'est pas dans la liste"`
	nullable      string `operator:"" func:"Nullable" message:"ne peut pas être vide"`
	between       string `operator:":" func:"Between" message:"doit être compris entre %v et %v"`
	date          string `operator:"" func:"Date" message:"n'est pas une date valide"`
	required      string `operator:"" func:"Required" message:"est requis"`
	required_with string `operator:"." func:"RequiredWith" message:"est requis avec le champ %v"`
}

func ApplyRule(rule string, nameAttribute string, value reflect.Value, formValue reflect.Value) []string {

	var errArray []string

	for _, v := range strings.Split(rule, "|") {
		structValue := reflect.ValueOf(RuleType{})
		structType := structValue.Type()
		innerRule, valueRule := getKeyRule(v)
		fieldType, exists := structType.FieldByName(innerRule)

		if exists {
			data := common.Parameter{
				Data:          value,
				ValueRule:     valueRule,
				AttributeName: nameAttribute,
			}

			if innerRule == "required_with" {
				data.StructRequest = formValue
				data.FieldType = fieldType
			}

			errFunc := callFunction(data, fieldType.Tag.Get("func"))

			for _, j := range errFunc {
				if !j.IsNil() {
					err := generateMessageErrorRule(nameAttribute, convertReflectValueToStringSlice(j), fieldType.Tag.Get("message"))
					fmt.Println(err)
					errArray = append(errArray, err)
				}
			}

		}

	}

	if nameAttribute == "id" {
		idDecode := helpers.DecodeId(value.Interface().(string))

		if idDecode == "0" {
			errArray = append(errArray, "L'id n'est pas valide")
		}

		value.SetString(idDecode)
	}

	return errArray
}

func ToArray(p common.Parameter) []reflect.Value {
	return []reflect.Value{reflect.ValueOf(p)}
}

func callFunction(fn common.Parameter, name string) []reflect.Value {
	method := reflect.ValueOf(&helpers.FuncCall{}).MethodByName(name)

	if method.IsValid() {
		return method.Call(ToArray(fn))
	}

	return nil
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

func generateMessageErrorRule(attribute string, value []string, message string) string {
	return fmt.Sprintf("Le champ << %s >> %s", attribute, replaceValueInMessage(message, value))
}

func replaceValueInMessage(message string, values []string) string {
	if strings.Contains(message, "%v") {
		for _, value := range values {
			message = strings.Replace(message, "%v", value, 1)
		}
		return message
	}

	return message
}

func ValuesToString(values []reflect.Value) string {
	var result string
	
	for _, v := range values {
		str := fmt.Sprintf("%v", v.Interface())
		result += str + "\n"
	}
	return result
}

func convertReflectValueToStringSlice(value reflect.Value) []string {
	if value.Kind() != reflect.Array && value.Kind() != reflect.Slice {
		return nil
	}

	sliceString := make([]string, value.Len())

	for i := 0; i < value.Len(); i++ {
		sliceString[i] = fmt.Sprintf("%v", value.Index(i))
	}

	return sliceString
}
