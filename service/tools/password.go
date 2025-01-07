package tools

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(pwd string) (s string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		err = errors.New("[repository][tool][SaltAndHash] while Generate From Password")
		log.Print(err)
		return s, Wrap(err)
	}

	return string(bytes), err
}
