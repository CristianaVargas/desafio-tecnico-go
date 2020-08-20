package main

import (
	"encoding/json"
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

// Claims will be encoded to a JWT
type Claims struct {
	AccountID int
	jwt.StandardClaims
}

// AuthenticateUser authenticate user
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var LoggedUser Login
	json.NewDecoder(r.Body).Decode(&LoggedUser)
	loggedAccount := FindAccountByCpf(LoggedUser.Cpf)
	if (Account{} != loggedAccount) {
		//not empty
		if CheckPasswordHash(LoggedUser.Secret, loggedAccount.Secret) {

			expirationTime := time.Now().Add(time.Minute * 30)
			signingKey = []byte(LoggedUser.Secret)

			claims := &Claims{
				AccountID: loggedAccount.ID,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			stringToken, err := token.SignedString([]byte(signingKey))

			if err != nil {
				w.WriteHeader(400)
				w.Write([]byte("Error generating token"))
				return
			}

			validToken = stringToken

			w.Write([]byte("Login Successful"))
			return
		}
	}
	w.WriteHeader(404)
	w.Write([]byte("Account or password not found"))

}

//IsAuthorized checks if user is authenticated
func IsAuthorized(endpoint http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		headToken := r.Header.Get("Token")
		// Initialize a new instance of `Claims`
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(headToken, claims, func(token *jwt.Token) (interface{}, error) {
			return signingKey, nil
		})

		if err != nil {
			w.WriteHeader(400)
			w.Write([]byte("User not authorized"))
			return
		}

		if !tkn.Valid {
			w.WriteHeader(400)
			w.Write([]byte("Token not valid"))
			return
		}

		endpoint.ServeHTTP(w, r)

	})
}

// AddTokenHeader adds token in the requisition's header
func AddTokenHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Token", validToken)
		next.ServeHTTP(w, r)
	})
}
