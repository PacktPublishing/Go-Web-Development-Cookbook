package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func updateDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	vals := r.URL.Query()
	name, ok := vals["name"]

	if ok {
		employeeId, err := strconv.Atoi(id)
		if err != nil {
			log.Print("error converting string id to int :: ", err)
			return
		}
		log.Print("going to update document in database for id :: ", id)
		collection := session.DB("mydb").C("employee")
		var changeInfo *mgo.ChangeInfo
		changeInfo, err = collection.Upsert(bson.M{"id": employeeId}, &Employee{employeeId, name[0]})
		if err != nil {
			log.Print("error occurred while updating record in database :: ", err)
			return
		}
		fmt.Fprintf(w, "Number of documents updated in database are :: %d", changeInfo.Updated)
	} else {
		fmt.Fprintf(w, "Error occurred while updating document in database for id :: %s", id)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/employee/update/{id}", updateDocument).Methods("PUT")
	defer session.Close()
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
