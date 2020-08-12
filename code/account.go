package main

import (
	"net/http"
	"time"
)

// Account is a struct that represents the accounts
type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Cpf       int       `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   float32   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

// GetAllAccounts get a list of all accounts
func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
}

// GetBalance get the balance of an account
func GetBalance(w http.ResponseWriter, r *http.Request) {

}

// CreateAccount create a new account
func CreateAccount(w http.ResponseWriter, r *http.Request) {

}
