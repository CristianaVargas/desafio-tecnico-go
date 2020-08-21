package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Account is a struct that represents the accounts
type Account struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Cpf       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   float32   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

// GetAllAccounts get a list of all accounts
func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Accounts)
}

// GetBalance get the balance of an account
func GetBalance(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["account_id"]

	tknAccountID := GetAccountIDFromToken(r.Header.Get("Token"))

	if tknAccountID == "" {
		w.WriteHeader(400)
		w.Write([]byte("Could not find Accound ID"))
		return
	}

	accountID, err := strconv.Atoi(tknAccountID)

	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	if id != accountID {
		w.WriteHeader(404)
		w.Write([]byte("Account not found"))
		return
	}

	s := fmt.Sprintf("%f", Accounts[id].Balance)
	w.Write([]byte("Your balance is: " + s))
}

// CreateAccount create a new account
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var newAccount Account
	json.NewDecoder(r.Body).Decode(&newAccount)
	checkCPF := FindAccountByCpf(newAccount.Cpf)
	if (checkCPF != Account{}) {
		w.WriteHeader(400)
		w.Write([]byte("CPF is already used"))
		return
	}
	newAccount.CreatedAt = time.Now()
	id := len(Accounts)
	newAccount.ID = id
	hash, err := HashPassword(newAccount.Secret)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Account could not be created. "))
		return
	}
	newAccount.Secret = hash
	Accounts = append(Accounts, newAccount)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Accounts)
}
