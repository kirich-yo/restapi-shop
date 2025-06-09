package req

import (
	"net/http"
)

func HandleBody[T any](r *http.Request) (*T, error) {
	body, err := DecodeDefault[T](r)
	if err != nil {
		return nil, err
	}

	err = IsValid[T](*body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
