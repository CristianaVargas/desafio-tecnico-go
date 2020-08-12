package main

import (
	"net/http"
	"time"
)

// Transfer is a struct that represents a transfer
type Transfer struct {
	ID                   string    `json:"id"`
	AccountOriginID      string    `json:"account_origin_id"`
	AccountDestinationID int       `json:"account_destination_id"`
	Amount               float32   `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

// GetTransfers get a list of all transfers
func GetTransfers(w http.ResponseWriter, r *http.Request) {

}

// MakeTransfer make a new transfer
func MakeTransfer(w http.ResponseWriter, r *http.Request) {

}
