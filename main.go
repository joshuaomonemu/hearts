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
}

//Renders Sign-Up Page

func signup(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/pages/sign-up.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}

//Rendering Sign-In Page

func sigin(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/pages/sign-in.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}

//Rendering Error Page

func errorPage(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/pages/404-page.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, nil)
}

//Rendering Home page after user login
func home(w http.ResponseWriter, r *http.Request) {
	tpl, _ := template.ParseFiles("templates/main.gohtml")
	err := tpl.Execute(w, nil)
	if err != nil {
		return
	}
}

func main() {
	setupRoutes()
	http.ListenAndServe(":2020", nil)
}
