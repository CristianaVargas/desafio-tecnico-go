package main

import (
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword convert password to hash
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash check if password matches with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// FindAccountByID find an account using ID
func FindAccountByID(accountID int) Account {
	var accountNotFount Account
	for _, account := range Accounts {
		if account.ID == accountID {
			return account
		}
	}
	return accountNotFount
}

// FindAccountByCpf find an account using CPF
func FindAccountByCpf(accountCpf string) Account {
	var accountNotFount Account
	for _, account := range Accounts {
		if account.Cpf == accountCpf {
			return account
		}
	}
	return accountNotFount
}

//GetAccountIDFromToken gets UserId from Claims
func GetAccountIDFromToken(token string) string {

	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if err != nil {
		return ""
	}

	return strconv.Itoa(claims.AccountID)
}
