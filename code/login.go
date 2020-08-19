package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var signingKey []byte

var validToken string

// Login is a struct that represents the login data
type Login struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

// AuthenticateUser autenticate user
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var LoggedUser Login
	json.NewDecoder(r.Body).Decode(&LoggedUser)
	for _, account := range Accounts {
		if account.Cpf == LoggedUser.Cpf {
			//Found account!
			if account.Secret == LoggedUser.Secret {

				signingKey = []byte(account.Secret)

				token, err := GenerateJWT()
				validToken = token
				if err != nil {
					fmt.Println("Error generating token string")
				}
				w.Write([]byte("Login Successful"))
				return
			}
		}
	}
	w.WriteHeader(404)
	w.Write([]byte("Account not found"))

}

//IsAuthorized checks if user is authenticated
func IsAuthorized(endpoint http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Token")
		if headerToken != "" {
			token, err := jwt.Parse(headerToken, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return signingKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				endpoint.ServeHTTP(w, r)
			}

		} else {
			w.WriteHeader(404)
			w.Write([]byte("User not authorized"))
		}
	})
}

//GenerateJWT to initialize authentication
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString([]byte(signingKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AddTokenHeader adds token in the requisition's header
func AddTokenHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Token", validToken)
		next.ServeHTTP(w, r)
	})
}
