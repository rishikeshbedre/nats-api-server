package util

import (
	//"fmt"

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
		//fmt.Println("Error producing bcrypt hash: ", err)
		return "", err
	}
	//fmt.Println("bcrypt hash: ", string(cb))
	return string(cb), err
}
