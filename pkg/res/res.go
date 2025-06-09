package res

import (
	"net/http"
	"strings"
	"encoding/json"
	"encoding/xml"
)

func WriteJSON(w http.ResponseWriter, code int, content any) {
	w.Header().Set("Content-Type", "application/json; encoding=utf-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(content)
}

func WriteXML(w http.ResponseWriter, code int, content any) {
	w.Header().Set("Content-Type", "application/xml; encoding=utf-8")
	w.WriteHeader(code)
	xml.NewEncoder(w).Encode(content)
}

func WriteDefault(w http.ResponseWriter, code int, content any, header http.Header) {
	accept_type := header.Get("Accept")
	if strings.HasPrefix(accept_type, "application/json") {
		WriteJSON(w, code, content)
		return
	}
	if strings.HasPrefix(accept_type, "application/xml") {
		WriteXML(w, code, content)
		return
	}
}
