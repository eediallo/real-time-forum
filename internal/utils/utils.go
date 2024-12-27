package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	//why does the CompareHashAndPassword in bcrypt package in golang return an err? and when?
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
