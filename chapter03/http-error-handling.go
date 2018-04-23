package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

type NameNotFoundError struct {
	Code int
	Err  error
}

func (nameNotFoundError NameNotFoundError) Error() string {
	return nameNotFoundError.Err.Error()
}

type WrapperHandler func(http.ResponseWriter, *http.Request) error

func (wrapperHandler WrapperHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := wrapperHandler(w, r)
	if err != nil {
		switch e := err.(type) {
		case NameNotFoundError:
			log.Printf("HTTP %s - %d", e.Err, e.Code)
			http.Error(w, e.Err.Error(), e.Code)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}

func getName(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	name := vars["name"]
	if strings.EqualFold(name, "foo") {
		fmt.Fprintf(w, "Hello "+name)
		return nil
	} else {
		return NameNotFoundError{500, errors.New("Name Not Found")}
	}
}

func main() {
	router := mux.NewRouter()
	router.Handle("/employee/get/{name}", WrapperHandler(getName)).Methods("GET")

	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
