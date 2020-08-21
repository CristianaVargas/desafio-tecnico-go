package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

// Transfer is a struct that represents a transfer
type Transfer struct {
	ID                   int       `json:"id"`
	AccountOriginID      int       `json:"account_origin_id"`
	AccountDestinationID int       `json:"account_destination_id"`
	Amount               float32   `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

// GetTransfers get a list of all transfers
func GetTransfers(w http.ResponseWriter, r *http.Request) {
	var accountTransfers []Transfer

	tknAccountID := GetAccountIDFromToken(r.Header.Get("Token"))
	if tknAccountID == "" {
		w.WriteHeader(404)
		w.Write([]byte("Origin ID not found"))
		return
	}
	accountID, err := strconv.Atoi(tknAccountID)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	totalTransfers := len(Transfers)
	for i := 0; i < totalTransfers; i++ {
		if Transfers[i].AccountOriginID == accountID {
			accountTransfers = append(accountTransfers, Transfers[i])
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(accountTransfers)

}

// MakeTransfer make a new transfer
func MakeTransfer(w http.ResponseWriter, r *http.Request) {
	var newTransfer Transfer
	tknAccountID := GetAccountIDFromToken(r.Header.Get("Token"))
	if tknAccountID == "" {
		w.WriteHeader(404)
		w.Write([]byte("Origin ID not found"))
		return
	}
	accountOriginID, err := strconv.Atoi(tknAccountID)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	originAccount := FindAccountByID(accountOriginID)

	json.NewDecoder(r.Body).Decode(&newTransfer)

	destinationAccount := FindAccountByID(newTransfer.AccountDestinationID)

	if (destinationAccount == Account{}) {
		w.WriteHeader(404)
		w.Write([]byte("Destination account not found"))
		return
	}

	if originAccount.Balance < newTransfer.Amount {
		w.WriteHeader(400)
		w.Write([]byte("You don't have enough balance"))
		return
	}
	originAccount.Balance -= newTransfer.Amount
	destinationAccount.Balance += newTransfer.Amount

	id := len(Transfers)
	newTransfer.ID = id
	newTransfer.CreatedAt = time.Now()

	Transfers = append(Transfers, newTransfer)
	Accounts[originAccount.ID] = originAccount
	Accounts[destinationAccount.ID] = destinationAccount

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Transfers)
	w.Write([]byte("Transference succesfull"))
	return

}
