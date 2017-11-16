package irc

import (
	"encoding/base64"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	EmptyPasswordError = errors.New("empty password")
)

func DecodePassword(encoded string) (decoded []byte, err error) {
	if encoded == "" {
		err = EmptyPasswordError
		return
	}
	decoded, err = base64.StdEncoding.DecodeString(encoded)
	return
}

func ComparePassword(hash, password []byte) error {
	return bcrypt.CompareHashAndPassword(hash, password)
}
