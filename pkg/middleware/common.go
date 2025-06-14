package middleware

import (
	"net/http"
)

type ResponseWriterWrapper struct {
	http.ResponseWriter
	StatusCode int
}

func NewResponseWriterWrapper(w http.ResponseWriter) *ResponseWriterWrapper {
	return &ResponseWriterWrapper{
		ResponseWriter: w,
	}
}

func (ww *ResponseWriterWrapper) WriteHeader(code int) {
	ww.ResponseWriter.WriteHeader(code)
	ww.StatusCode = code
}
