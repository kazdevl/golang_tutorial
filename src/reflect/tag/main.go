package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
)

const TagName = "custom"

const (
	TypeMin     = "min"
	TypeMax     = "max"
	TypeDefault = "default"
)

type Human struct {
	Weight int    `custom:"min=20"`
	Height int    `custom:"max=200"`
	Name   string `custom:"default=sample"`
}

func main() {
	// error in weifht
	h := &Human{
		Weight: 10,
		Height: 180,
		Name:   "Hoge",
	}
	if err := Validate(h); err != nil {
		log.Println(err)
	}
	// error in height
	h = &Human{
		Weight: 60,
		Height: 210,
		Name:   "Hoge",
	}
	if err := Validate(h); err != nil {
		log.Println(err)
	}
	// update name
	h = &Human{
		Weight: 60,
		Height: 200,
		Name:   "",
	}
	if err := Validate(h); err != nil {
		log.Println(err)
	}
	fmt.Println(h)
	// not update name
	h = &Human{
		Weight: 60,
		Height: 210,
		Name:   "Hoge",
	}
	if err := Validate(h); err != nil {
		log.Println(err)
	}
	fmt.Println(h)
	fmt.Println("finish")
}

func Validate(data interface{}) error {
	typ := reflect.TypeOf(data)
	val := reflect.ValueOf(data)
	fmt.Printf("name: %s\n", typ.Elem().Name())
	if typ.Elem().Name() != "Human" {
		return errors.New("unexpected struct")
	}
	switch typ.Kind() {
	case reflect.Ptr:
		for i := 0; i < typ.Elem().NumField(); i++ {
			fieldTyp := typ.Elem().Field(i)
			fieldVal := val.Elem().Field(i)
			fmt.Printf("fieldTyp: %+v\n", fieldTyp)
			fmt.Printf("fieldVal: %+v\n", fieldVal)
			if err := validateField(fieldTyp, fieldVal); err != nil {
				return err
			}
		}
	default:
		return errors.New("unexpected Kind")
	}
	return nil
}

func validateField(typ reflect.StructField, val reflect.Value) error {
	tagVal, ok := typ.Tag.Lookup(TagName)
	if !ok {
		return errors.New("not found target tag name")
	}
	fmt.Println(tagVal)
	tagvalInfo := strings.Split(tagVal, "=")
	if len(tagvalInfo) != 2 {
		return errors.New("not correct tag format")
	}
	switch tagvalInfo[0] {
	case TypeMin:
		if typ.Type.Kind() != reflect.Int {
			return errors.New("field type is not int")
		}
		min, _ := strconv.Atoi(tagvalInfo[1])
		if int(val.Int()) < min {
			return fmt.Errorf("min: %d, val: %d\n", min, int(val.Int()))
		}
	case TypeMax:
		if typ.Type.Kind() != reflect.Int {
			return errors.New("field type is not int")
		}
		max, _ := strconv.Atoi(tagvalInfo[1])
		if int(val.Int()) > max {
			return fmt.Errorf("max: %d, val: %d\n", max, int(val.Int()))
		}
	case TypeDefault:
		if typ.Type.Kind() != reflect.String {
			return errors.New("field type is not string")
		}
		if len(val.String()) == 0 {
			val.SetString(tagvalInfo[1])
		}
	}
	return nil
}
