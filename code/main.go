package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Accounts is the list of accounts
var Accounts []Account = []Account{}

//Transfers is the list of transfers of an account
var Transfers []Transfer = []Transfer{}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/accounts", GetAllAccounts).Methods("GET")

	router.HandleFunc("/accounts/{account_id}/balance", GetBalance).Methods("GET")

	router.HandleFunc("/accounts", CreateAccount).Methods("POST")

	router.HandleFunc("/login", VerifyLogin).Methods("POST")

	router.HandleFunc("/transfers", GetTransfers).Methods("GET")

	router.HandleFunc("/transfers", MakeTransfer).Methods("POST")

	http.ListenAndServe(":5000", router)
}
