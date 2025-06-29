package form

import (
	"reflect"
	"fmt"
	"strings"
)

func Encode(data interface{}) (string, error) {
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Pointer {
		return "", ErrNotPointerToStruct
	}

	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return "", ErrNotPointerToStruct
	}

	v := reflect.ValueOf(data).Elem()
	formValues := make([]string, 0)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("form")
		if tag != "" {
			fieldValue := v.Field(i)
			formValues = append(formValues, fmt.Sprintf("%s=%s", tag, fieldValue.String()))
		}
	}

	return strings.Join(formValues, "&"), nil
}

func Decode(data interface{}, encodedForm string) error {
	t := reflect.TypeOf(data)
	if t.Kind() != reflect.Pointer {
		return ErrNotPointerToStruct
	}

	t = t.Elem()
	if t.Kind() != reflect.Struct {
		return ErrNotPointerToStruct
	}

	v := reflect.ValueOf(data).Elem()
	formValues := strings.Split(encodedForm, "&")

	for _, fv := range formValues {
		equalIdx := strings.IndexByte(fv, '=')
		if equalIdx == -1 {
			return ErrInvalidFormat
		}

		key := fv[:equalIdx]
		value := fv[equalIdx+1:]

		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			tag := field.Tag.Get("form")
			if tag == key {
				v.Field(i).SetString(value)
			}
		}
	}

	return nil
}
