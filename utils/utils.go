package utils

import (
	"github.com/gofrs/uuid/v5"
)

func NewUUIDv7() string {
	u, err := uuid.NewV7()
	if err != nil {
		panic(err) // Em produção, trate o erro de forma adequada
	}
	return u.String()
}
