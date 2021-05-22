package main

import "net/http"

func setupRoutes() {
	//Help
	//Rendering Server Files
	http.Handle("/", http.FileServer(http.Dir("./templates")))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//Rendering Server Pages
	http.HandleFunc("/index", signup)
	http.HandleFunc("/signin", sigin)
	http.HandleFunc("/home", home)
	http.HandleFunc("/error", errorPage)
	//Handling Actions
	http.HandleFunc("/signinAction", signinAction)
	http.HandleFunc("/signupAction", signupAction)
	//Start Server
}
