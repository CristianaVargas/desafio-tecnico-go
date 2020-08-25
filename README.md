# Desafio Técnico em Go

API de transferencia entre contas Internas de um banco digital.

### Install
```
go build
````
### Run app
```
./code
```

## REST API
As operações são descritas abaixo.
## Get list of accounts
### Request
```
GET http://localhost:5000/accounts
```
### Response
```
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Fri, 21 Aug 2020 02:59:09 GMT
< Content-Length: 356

[
  {
    "id": 0,
    "name": "John Doe",
    "cpf": "123-56",
    "secret": "$2a$14$Hf5WysqovK4tdrbNyDUjv.wPN5afSH6Kb2/HvPggFKv6gcAyMW/Oi",
    "balance": 300,
    "created_at": "2020-08-20T23:58:33.8273116-03:00"
  }
]
```
## Create an account 
### Request
```
POST http://localhost:5000/accounts
```
```
JSON Body
{
	"name": "John Doe",
	"cpf": "456-56",
	"secret": "password",
	"balance": 300.00
}
```
### Response
```
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Fri, 21 Aug 2020 02:58:48 GMT
< Content-Length: 356

[
  {
      "id": 0,
      "name": "John Doe",
      "cpf": "123-56",
      "secret": "$2a$14$Hf5WysqovK4tdrbNyDUjv.wPN5afSH6Kb2/HvPggFKv6gcAyMW/Oi",
      "balance": 300,
      "created_at": "2020-08-20T23:58:33.8273116-03:00"
    }
]
```
## Atuhenticate user
### Request
```
POST http://localhost:5000/login
```
```
JSON Body
{
	"cpf": "123-56",
	"secret": "password"
}
```
### Response
```
< HTTP/1.1 200 OK
< Date: Fri, 21 Aug 2020 02:59:15 GMT
< Content-Length: 16
< Content-Type: text/plain; charset=utf-8

Login Successful
```
## Get account balance
User must be logged in
### Request
```
GET http://localhost:5000/accounts/{account_id}/balance
```
### Response
```
< HTTP/1.1 200 OK
< Date: Thu, 20 Aug 2020 21:34:33 GMT
< Content-Length: 27
< Content-Type: text/plain; charset=utf-8

Your balance is: 300.000000
```
## Make a new transfer
User must be logged in
### Request
```
POST http://localhost:5000/transfers
```
```
JSON Body
{
	"account_destination_id": 1,
	"amount": 10.00
}
```
### Response
```
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Fri, 21 Aug 2020 02:59:24 GMT
< Content-Length: 23

Transference succesfull
```
## Get list of transfers for one user
User must be logged in
### Request
```
GET http://localhost:5000/transfers
```
### Response
```
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Fri, 21 Aug 2020 03:08:27 GMT
< Content-Length: 121

[
  {
    "id": 0,
    "account_origin_id": 0,
    "account_destination_id": 1,
    "amount": 10,
    "created_at": "2020-08-20T23:59:24.2667226-03:00"
  }
]
```
