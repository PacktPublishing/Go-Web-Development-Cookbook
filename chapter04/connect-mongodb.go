package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	mgo "gopkg.in/mgo.v2"
)

const (
	CONN_HOST    = "localhost"
	CONN_PORT    = "8080"
	MONGO_DB_URL = "127.0.0.1"
)

var session *mgo.Session
var connectionError error

func init() {
	session, connectionError = mgo.Dial(MONGO_DB_URL)
	if connectionError != nil {
		log.Fatal("error connecting to database :: ", connectionError)
	}
	session.SetMode(mgo.Monotonic, true)
}

func getDbNames(w http.ResponseWriter, r *http.Request) {
	db, err := session.DatabaseNames()
	if err != nil {
		log.Print("error getting database names :: ", err)
		return
	}
	fmt.Fprintf(w, "Databases names are :: %s", strings.Join(db, ", "))
}

func main() {
	http.HandleFunc("/", getDbNames)
	defer session.Close()
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server :: ", err)
		return
	}
}
