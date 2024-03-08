package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HashPassword(password string) (string, error) {
	h := sha256.New()
	h.Write([]byte(password))
	resultHash := h.Sum(nil)
	resultString := hex.EncodeToString(resultHash)
	return resultString, nil
}

// _HashPassword return two diff hash for equal password
func _HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func _ComparePasswords(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytesPassword := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytesPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
