package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var connectionError error

const (
	CONN_PORT        = "8080"
	DRIVER_NAME      = "mysql"
	DATA_SOURCE_NAME = "root:my-pass@tcp(mysql-container:3306)/mysql"
)

func init() {
	db, connectionError = sql.Open(DRIVER_NAME, DATA_SOURCE_NAME)
	if connectionError != nil {
		log.Fatal("error connecting to database : ", connectionError)
	}
}

func getDBInfo(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT SUBSTRING_INDEX(USER(), '@', -1) AS ip,  @@hostname as hostname, @@port as port, DATABASE() as current_database;")
	if err != nil {
		log.Print("error executing database query : ", err)
		return
	}
	var buffer bytes.Buffer

	for rows.Next() {
		var ip string
		var hostname string
		var port string
		var current_database string
		err = rows.Scan(&ip, &hostname, &port, &current_database)
		buffer.WriteString("IP :: " + ip + " | HostName :: " + hostname + " | Port :: " + port + " | Current Database :: " + current_database)
	}

	fmt.Fprintf(w, buffer.String())
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", getDBInfo).Methods("GET")
	defer db.Close()
	err := http.ListenAndServe(":"+CONN_PORT, router)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}

}
