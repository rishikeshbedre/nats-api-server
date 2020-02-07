package util

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	//DefaultCost is set to 11, you can raise this value.
	DefaultCost = 11
)

//GenHashPassword function generates bcrypt password from plain text
func GenHashPassword(pw string) (string, error) {
	cb, err := bcrypt.GenerateFromPassword([]byte(pw), DefaultCost)
	if err != nil {
		return "", err
	}
	return string(cb), err
}
