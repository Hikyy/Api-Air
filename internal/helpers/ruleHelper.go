package helpers

import (
	"App/internal/common"
	"errors"
	"fmt"
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

func (f *FuncCall) Email(attribute common.Parameter) error {
	regex := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	if regex.FindString(attribute.Data.String()) == "" {
		return errors.New("test")
	}

	return nil
}

func (f *FuncCall) String(attribute common.Parameter) error {
	if reflect.TypeOf(attribute.Data).Kind() != reflect.String {
		return errors.New("test")
	}
	return nil
}

func (f *FuncCall) Integer(attribute common.Parameter) error {
	if reflect.TypeOf(attribute.Data).Kind() != reflect.Int {
		return errors.New("int not found")
	}
	return nil
}

func (f *FuncCall) Nullable(attribute common.Parameter) error {
	fmt.Println("nullable")
	if attribute.Data.String() == "" {
		return errors.New("test")
	}
	return nil
}

func (f *FuncCall) Required(attribute common.Parameter) error {
	if attribute.Data.String() == "" || reflect.TypeOf(attribute.Data) == nil {
		return errors.New("hihi")
	}
	return nil
}

func (f *FuncCall) Max(attribute common.Parameter) error {
	numberMax, err := strconv.Atoi(attribute.ValueRule)
	if err != nil {
		return errors.New("test")
	}

	if len(attribute.Data.String()) > numberMax {
		return errors.New("test")
	}
	return nil
}

func (f *FuncCall) Min(attribute common.Parameter) error {
	numberMin, err := strconv.Atoi(attribute.ValueRule)
	if err != nil {
		return errors.New("test")
	}

	if len(attribute.Data.String()) > numberMin {
		return errors.New("test")
	}
	return nil
}
