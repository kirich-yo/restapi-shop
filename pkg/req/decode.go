package req

import (
	"net/http"
	"strings"
	"encoding/json"
	"encoding/xml"
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
	return nil, ErrUnsupportedContentType
}
