package helpers

import (
	"App/internal/common"
	"fmt"
	"strings"

	"reflect"
	"regexp"
	"strconv"
)

type FuncCall struct{}

type FuncCallType interface {
	Email(attribute common.Parameter) error
	Integer(value common.Parameter) error
	String(attribute common.Parameter) error
	Required(attribute common.Parameter) error
	Min(attribute common.Parameter) error
	Max(attribute common.Parameter) error
	Unique(attribute common.Parameter) error
	In(attribute common.Parameter) error
	Nullable(attribute common.Parameter) error
	Between(attribute common.Parameter) error
	Date(attribute common.Parameter) error
	RequiredWith(attribute common.Parameter) error
}

func (f *FuncCall) Email(attribute common.Parameter) []string {
	regex := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	if regex.FindString(attribute.Data.String()) == "" {
		return []string{attribute.AttributeName}
	}

	return nil
}

func (f *FuncCall) String(attribute common.Parameter) []string {
	if reflect.TypeOf(attribute.Data).Kind() != reflect.String {
		return []string{attribute.AttributeName}
	}
	return nil
}

func (f *FuncCall) Integer(attribute common.Parameter) []string {
	if reflect.TypeOf(attribute.Data).Kind() != reflect.Int {
		return []string{attribute.AttributeName}
	}
	return nil
}

func (f *FuncCall) Nullable(attribute common.Parameter) []string {
	if attribute.Data.String() == "" {
		return []string{attribute.AttributeName}
	}
	return nil
}

func (f *FuncCall) Required(attribute common.Parameter) []string {
	if attribute.Data.String() == "" || reflect.TypeOf(attribute.Data) == nil {
		return []string{attribute.AttributeName}
	}
	return nil
}

func (f *FuncCall) Max(attribute common.Parameter) []string {
	numberMax, err := strconv.Atoi(attribute.ValueRule)

	if err != nil {
		return []string{attribute.AttributeName}
	}

	if len(attribute.Data.String()) > numberMax {
		return []string{attribute.AttributeName}
	}
	return nil
}

func (f *FuncCall) Min(attribute common.Parameter) []int {
	numberMin, _ := strconv.Atoi(attribute.ValueRule)

	value := len(attribute.Data.String())

	if value > numberMin {
		return nil
	}

	return []int{numberMin}
}

func (f *FuncCall) Between(attribute common.Parameter) []int {
	attributeSplit := strings.Split(attribute.ValueRule, ",")

	numberMin, _ := strconv.Atoi(attributeSplit[0])
	numberMax, _ := strconv.Atoi(attributeSplit[1])

	data, _ := strconv.Atoi(attribute.Data.String())

	if numberMin < data && numberMax > data {
		return nil
	}

	return []int{numberMin, numberMax}
}

func (f *FuncCall) RequiredWith(attribute common.Parameter) []string {
	var sliceAttribute []string = strings.Split(attribute.ValueRule, attribute.FieldType.Tag.Get("operator"))

	err := checkFieldEmpty(sliceAttribute, attribute.StructRequest)
	if err != nil {
		return []string{attribute.ValueRule}
	}
	return nil
}

func checkFieldEmpty(value []string, structRequest reflect.Value) []string {
	l := len(value) - 1
	lastAttribute := toTitleCase(value[l])

	fieldValue := structRequest.FieldByName(lastAttribute)

	if fieldValue.IsValid() {
		if fieldValue.String() != "" {
			return nil
		} else {
			return value
		}
	}
	return []string{"field not found"}
}

func toTitleCase(value string) string {
	words := strings.Split(value, "_")

	for i := 0; i < len(words); i++ {
		words[i] = strings.Title(words[i])
	}
	camelCaseStr := strings.Join(words, "")

	return camelCaseStr
}

func DecodeId(id string) string {
	o := InitOptimus()
	numberMin, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Erreur lors de l'encodage:", err)
	}

	newId := o.Decode(uint64(numberMin))

	return strconv.FormatUint(newId, 10)
}
