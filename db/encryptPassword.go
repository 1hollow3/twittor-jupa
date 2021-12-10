package db

import "golang.org/x/crypto/bcrypt"


func EncryptPassword(pass string)(string, error){
	cost := 6
	by, err := bcrypt.GenerateFromPassword([]byte(pass), cost )
	return string(by), err
}