package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Accounts is the list of accounts
var Accounts []Account = []Account{}

//Transfers is the list of transfers of an account
var Transfers []Transfer = []Transfer{}

//LoggedUser from current user
var LoggedUser Login

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/login", AuthenticateUser).Methods("POST")

	router.HandleFunc("/accounts", GetAllAccounts).Methods("GET")

	router.HandleFunc("/accounts", CreateAccount).Methods("POST")

	router.Handle("/accounts/{account_id}/balance", AddTokenHeader(IsAuthorized(http.HandlerFunc(GetBalance)))).Methods("GET")

	router.Handle("/transfers", AddTokenHeader(IsAuthorized(http.HandlerFunc(GetTransfers)))).Methods("GET")

	router.Handle("/transfers", AddTokenHeader(IsAuthorized(http.HandlerFunc(MakeTransfer)))).Methods("POST")

	http.ListenAndServe(":5000", router)
}
