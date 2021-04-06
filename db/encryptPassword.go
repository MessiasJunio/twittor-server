package db

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(passwd string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd), cost)
	return string(bytes), err
}
