package main

import (
	"log"
	"net/http"
	"text/template"
)

//Struct For User Details
type person struct {
	Person userInfo `json:"person"`
}
type userInfo struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Blocked   string `json:"blocked"`
	Location string `json:"location"`
	Email string `json:"email"`
}

//Renders Sign-Up Page

func signup(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("app/sign-up.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}

//Rendering Sign-In Page

func sigin(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("app/sign-in.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}

//Rendering Error Page

func errorPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("app/404-page.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}
func main() {
	setupRoutes()
	http.ListenAndServe(":2020", nil)
}
