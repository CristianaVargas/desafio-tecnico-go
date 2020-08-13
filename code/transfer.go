package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Transfer is a struct that represents a transfer
type Transfer struct {
	ID                   int       `json:"id"` //obrigatorios os OrigenID, DestinationID e Amount
	AccountOriginID      int       `json:"account_origin_id"`
	AccountDestinationID int       `json:"account_destination_id"`
	Amount               float32   `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

// GetTransfers get a list of all transfers
func GetTransfers(w http.ResponseWriter, r *http.Request) {
	var accountTransfers []Transfer
	var idParam string = mux.Vars(r)["account_origin_id"]
	accountID, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	//error checking
	if accountID >= len(Transfers) {
		w.WriteHeader(404)
		w.Write([]byte("No transfer found for specified account"))
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
	json.NewDecoder(r.Body).Decode(&newTransfer)
	for _, i := range Accounts {
		if i.ID == newTransfer.AccountDestinationID {
			// Found!
			var originAccount = Accounts[newTransfer.AccountOriginID]
			var destinationAccount = Accounts[newTransfer.AccountDestinationID]

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

			w.Write([]byte("Transference succesfull"))
			return
		}
	}

	w.WriteHeader(404)
	w.Write([]byte("ID of destination account could not be found"))
}
