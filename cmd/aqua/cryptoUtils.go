package main

import "golang.org/x/crypto/bcrypt"

const COST  = bcrypt.DefaultCost


func getHash(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), COST)
	if err != nil {
		EXCEPTION(err)
	}
	return string(hash)
}