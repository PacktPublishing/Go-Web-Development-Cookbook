package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
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

func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API is up and running"))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(employees)
}

func getToken(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/get-token", getToken).Methods("GET")
	router.HandleFunc("/employees", getEmployees).Methods("GET")
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, handlers.LoggingHandler(os.Stdout, router))
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
