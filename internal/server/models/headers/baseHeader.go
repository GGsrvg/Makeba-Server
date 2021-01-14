package headers

import (
	"errors"
	"net/http"
)

type BaseHeader struct {
	Token string
}

func DetectBaseHeader(header *http.Header) (*BaseHeader, error) {
	token := header.Get("Token")
	if token == "" {
		return nil, errors.New("Exist token")
	}
	return &BaseHeader{Token: token}, nil
}
