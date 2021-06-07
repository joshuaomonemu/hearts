package main

import (
	"net/http"
)

func setupRoutes() {
	//Help
	//Rendering Server Files
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./templates/assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//Rendering Server Pages
	http.HandleFunc("/", signup)
	http.HandleFunc("/signin", sigin)
	http.HandleFunc("/home", home)
	http.HandleFunc("/error", errorPage)
	http.HandleFunc("/main", renderUser)
	//Handling Actions
	http.HandleFunc("/signinAction", signinAction)
	http.HandleFunc("/signupAction", signupAction)
	http.HandleFunc("/logout", logOut)

	// 404 page
	//http.HandleFunc("/", notFound)

}
