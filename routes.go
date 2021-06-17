package main

import (
	"net/http"
)

func setupRoutes() {
	//Help
	lostroutes := "*/.."
	//Rendering Server Files
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./app/assets"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())
	//Rendering Server Pages
	http.HandleFunc("/", signup)
	http.HandleFunc("/signin", sigin)
	http.HandleFunc("/error", errorPage)
	http.HandleFunc("/main", renderUser)
	http.HandleFunc("/ws", wsEndpoint)
	http.HandleFunc(lostroutes, lostPage)
	//Handling Actions
	http.HandleFunc("/signinAction", signinAction)
	http.HandleFunc("/signupAction", signupAction)
	http.HandleFunc("/logout", logOut)

	// 404 page
	//http.HandleFunc("/", notFound)

}
