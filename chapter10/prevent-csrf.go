package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
)

const (
	CONN_HOST          = "localhost"
	CONN_PORT          = "8443"
	HTTPS_CERTIFICATE  = "domain.crt"
	DOMAIN_PRIVATE_KEY = "domain.key"
)

var AUTH_KEY = []byte("authentication-key")

func signUp(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("sign-up.html")
	err := parsedTemplate.Execute(w, map[string]interface{}{
		csrf.TemplateTag: csrf.TemplateField(r),
	})
	if err != nil {
		log.Printf("Error occurred while executing the template : ", err)
		return
	}
}

func post(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Print("error occurred while parsing form ", err)
	}
	name := r.FormValue("name")
	fmt.Fprintf(w, "Hi %s", name)
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	muxRouter.HandleFunc("/signup", signUp)
	muxRouter.HandleFunc("/post", post)
	http.ListenAndServeTLS(CONN_HOST+":"+CONN_PORT, HTTPS_CERTIFICATE, DOMAIN_PRIVATE_KEY, csrf.Protect(AUTH_KEY)(muxRouter))
}
