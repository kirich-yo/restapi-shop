package review

import (
	"encoding/xml"
)

type ReviewCreateRequest struct {
	XMLName xml.Name `json:"-" xml:"review"`
}
