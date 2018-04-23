package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	CONN_HOST    = "localhost"
	CONN_PORT    = "8080"
	MONGO_DB_URL = "127.0.0.1"
)

var session *mgo.Session
var connectionError error

type Employee struct {
	Id   int    `json:"uid"`
	Name string `json:"name"`
}

func init() {
	session, connectionError = mgo.Dial(MONGO_DB_URL)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
	session.SetMode(mgo.Monotonic, true)
}

func deleteDocument(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	name, ok := vals["name"]
	if ok {
		log.Print("going to delete document in database for name :: ", name[0])
		collection := session.DB("mydb").C("employee")
		removeErr := collection.Remove(bson.M{"name": name[0]})
		if removeErr != nil {
			log.Print("error removing document from database :: ", removeErr)
			return
		}
		fmt.Fprintf(w, "Document with name %s is deleted from database", name[0])
	} else {
		fmt.Fprintf(w, "Error occurred while deleting document in database for name :: %s", name[0])
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee/delete", deleteDocument).Methods("DELETE")
	defer session.Close()
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
