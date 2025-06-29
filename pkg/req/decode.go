package req

import (
	"net/http"
	"strings"
	"encoding/json"
	"encoding/xml"
	"bufio"

	"restapi-sportshop/pkg/form"
)

func DecodeJSON[T any](r *http.Request) (*T, error) {
	var payload T

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func DecodeXML[T any](r *http.Request) (*T, error) {
	var payload T

	err := xml.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func DecodeForm[T any](r *http.Request) (*T, error) {
	var payload T

	rd := bufio.NewReader(r.Body)
	encodedForm, _ := rd.ReadString('\n')

	err := form.Decode(&payload, encodedForm)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}

func DecodeDefault[T any](r *http.Request) (*T, error) {
	content_type := r.Header.Get("Content-Type")
        if strings.HasPrefix(content_type, "application/json") {
		payload, err := DecodeJSON[T](r)
		if err != nil {
			return nil, err
		}
		return payload, nil
        }
        if strings.HasPrefix(content_type, "application/xml") {
		payload, err := DecodeXML[T](r)
		if err != nil {
			return nil, err
		}
		return payload, nil
        }
        if strings.HasPrefix(content_type, "application/x-www-form-urlencoded") {
		payload, err := DecodeForm[T](r)
		if err != nil {
			return nil, err
		}
		return payload, nil
        }
	return nil, ErrUnsupportedContentType
}
