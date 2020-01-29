package util

import (
	//"fmt"

	"golang.org/x/crypto/bcrypt"
)

const (
	// Make sure the password is reasonably long to generate enough entropy.
	//PasswordLength = 22
	// Common advice from the past couple of years suggests that 10 should be sufficient.
	// Up that a little, to 11. Feel free to raise this higher if this value from 2015 is
	// no longer appropriate. Min is bcrypt.MinCost, Max is bcrypt.MaxCost.
	DefaultCost = 11
)

func GenHashPassword(pw string) (string, error) {

	cb, err := bcrypt.GenerateFromPassword([]byte(pw), DefaultCost)
	if err != nil {
		//fmt.Println("Error producing bcrypt hash: ", err)
		return "", err
	}
	//fmt.Println("bcrypt hash: ", string(cb))
	return string(cb), err
}
