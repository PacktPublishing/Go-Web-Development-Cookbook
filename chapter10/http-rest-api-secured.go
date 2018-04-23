package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST             = "localhost"
	CONN_PORT             = "8080"
	CLAIM_ISSUER          = "Packt"
	CLAIM_EXPIRY_IN_HOURS = 24
)

type Employee struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

var employees []Employee

func init() {
	employees = Employees{
		Employee{Id: 1, FirstName: "Foo", LastName: "Bar"},
		Employee{Id: 2, FirstName: "Baz", LastName: "Qux"},
	}
}

var signature = []byte("secret")

var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return signature, nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

func getToken(w http.ResponseWriter, r *http.Request) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * CLAIM_EXPIRY_IN_HOURS).Unix(),
		Issuer:    CLAIM_ISSUER,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString(signature)
	w.Write([]byte(tokenString))
}

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/status", getStatus).Methods("GET")
	muxRouter.HandleFunc("/get-token", getToken).Methods("GET")
	muxRouter.Handle("/employees", jwtMiddleware.Handler(http.HandlerFunc(getEmployees))).Methods("GET")
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, handlers.LoggingHandler(os.Stdout, muxRouter))
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
