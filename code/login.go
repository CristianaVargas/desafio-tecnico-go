package main

import "net/http"

// Login is a struct that represents the login data
type Login struct {
	Cpf    int    `json:"cpf"`
	Secret string `json:"secret"`
}

// VerifyLogin autenticate user
func VerifyLogin(w http.ResponseWriter, r *http.Request) {

}
