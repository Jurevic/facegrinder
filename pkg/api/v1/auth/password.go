package auth

import (
	"log"
	"golang.org/x/crypto/bcrypt"
)

func hashAndSalt(pwd []byte) []byte {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}

	return hash
}

func checkPasswordMatch(hash, password []byte) error {
	// Compares password with hash using bcrypt if passwords
	// do not match returns err
	return bcrypt.CompareHashAndPassword(hash, password)
}


