package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func getUserName(request *http.Request) (userName string) {
	cookie, err := request.Cookie("session")

	if err == nil {
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil {
			userName = cookieValue["username"]
		}
	}
	return userName
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string{
		"username": userName,
	}
	encoded, err := cookieHandler.Encode("session", value)

	if err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func login(response http.ResponseWriter, request *http.Request) {
	username := request.FormValue("username")
	password := request.FormValue("password")
	target := "/"
	if username != "" && password != "" {
		setSession(username, response)
		target = "/home"
	}
	http.Redirect(response, request, target, 302)
}

func logout(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/", 302)
}

func loginPage(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, _ := template.ParseFiles("templates/login-form.html")
	parsedTemplate.Execute(w, nil)
}

func homePage(response http.ResponseWriter, request *http.Request) {
	userName := getUserName(request)
	if userName != "" {
		data := map[string]interface{}{
			"userName": userName,
		}
		parsedTemplate, _ := template.ParseFiles("templates/home.html")
		parsedTemplate.Execute(response, data)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}

func main() {
	var router = mux.NewRouter()
	router.HandleFunc("/", loginPage)
	router.HandleFunc("/home", homePage)

	router.HandleFunc("/login", login).Methods("POST")
	router.HandleFunc("/logout", logout).Methods("POST")

	http.Handle("/", router)
	err := http.ListenAndServe(CONN_HOST+":"+CONN_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server : ", err)
		return
	}
}
