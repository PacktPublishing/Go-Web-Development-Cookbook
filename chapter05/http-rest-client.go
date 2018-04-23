package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	resty "gopkg.in/resty.v1"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8090"
)

const WEB_SERVICE_HOST string = "http://localhost:8080"

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	response, err := resty.R().Get(WEB_SERVICE_HOST + "/employees")
	if err != nil {
		log.Print("error getting data from the web service :: ", err)
		return
	}
	printOutput(response, err)
	fmt.Fprintf(w, response.String())
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	decodingErr := json.NewDecoder(r.Body).Decode(&employee)
	if decodingErr != nil {
		log.Print("error occurred while decoding employee data :: ", decodingErr)
		return
	}
	log.Printf("adding employee id :: %s with firstName as :: %s and lastName as :: %s ", employee.Id, employee.FirstName, employee.LastName)

	response, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(Employee{Id: employee.Id, FirstName: employee.FirstName, LastName: employee.LastName}).
		Post(WEB_SERVICE_HOST + "/employee/add")

	if err != nil {
		log.Print("error occurred while adding employee :: ", err)
		return
	}
	printOutput(response, err)
	fmt.Fprintf(w, response.String())
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	decodingErr := json.NewDecoder(r.Body).Decode(&employee)
	if decodingErr != nil {
		log.Print("error occurred while decoding employee data :: ", decodingErr)
		return
	}
	log.Printf("updating employee id :: %s with firstName as :: %s and lastName as :: %s ", employee.Id, employee.FirstName, employee.LastName)
	response, err := resty.R().
		SetBody(Employee{Id: employee.Id, FirstName: employee.FirstName, LastName: employee.LastName}).
		Put(WEB_SERVICE_HOST + "/employee/update")

	if err != nil {
		log.Print("error occurred while updating employee :: ", err)
		return
	}
	printOutput(response, err)
	fmt.Fprintf(w, response.String())
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	employee := Employee{}
	decodingErr := json.NewDecoder(r.Body).Decode(&employee)
	if decodingErr != nil {
		log.Print("error occurred while decoding employee data :: ", decodingErr)
		return
	}
	log.Printf("deleting employee id :: %s with firstName as :: %s and lastName as :: %s ", employee.Id, employee.FirstName, employee.LastName)
	response, err := resty.R().
		SetBody(Employee{Id: employee.Id, FirstName: employee.FirstName, LastName: employee.LastName}).
		Delete(WEB_SERVICE_HOST + "/employee/delete")

	if err != nil {
		log.Print("error occurred while deleting employee :: ", err)
		return
	}
	printOutput(response, err)
	fmt.Fprintf(w, response.String())
}

func printOutput(resp *resty.Response, err error) {
	log.Println(resp, err)
}

func main() {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/employees", getEmployees).Methods("GET")
	employee := router.PathPrefix("/employee").Subrouter()
	employee.HandleFunc("/add", addEmployee).Methods("POST")
	employee.HandleFunc("/update", updateEmployee).Methods("PUT")
	employee.HandleFunc("/delete", deleteEmployee).Methods("DELETE")

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
