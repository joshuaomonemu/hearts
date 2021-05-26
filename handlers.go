package main

import (
	helpers "https://github.com/joshuaomonemu/spades/utility/helper"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Handles Sign-in Request

func signinAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request", 400)
	}
	user := r.FormValue("username")
	password := r.FormValue("password")

	userFile := "user/" + user + ".gojson"
	reply, err := ioutil.ReadFile(userFile)

	if err != nil {
		http.Error(w, "Sorry couldnt read user files", 500)
	}

	//Map that stores values read from GOJSON
	var passer map[string]interface{}

	//Converting GOJSON values and storing in a MAP
	err = json.Unmarshal(reply, &passer)
	if err != nil {
		http.Error(w, "Couldn't read user file", 500)
	}

	//Check if username and password are valid
	userField := passer["person"].(map[string]interface{})["username"]
	passField := passer["person"].(map[string]interface{})["password"]

	//Check if user account  is blocked
	block := passer["person"].(map[string]interface{})["blocked"]
	accnt_block := fmt.Sprintf("%v", block)
	//What happens if account is blocked
	if accnt_block == "true" {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}

	//Conditional statments to ensure struct values are thesame
	if passField == password || userField == user {
		http.Redirect(w, r, "/home", http.StatusFound)
	} else {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, `<h1>Incorrect Password Or Username</h1>`)
	}
}
func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Thanks for loggin")
}

//

//Handles Sign-up Request
func signupAction(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		http.Error(w, "Bad Request", 404)
	}
	fname := r.FormValue("firstname")
	lname := r.FormValue("lastname")
	user := r.FormValue("username")
	pass := r.FormValue("password")

	//Sanitizing Form Values
	if fname == "" || lname == "" {
		http.Error(w, "Fill in empty fields", http.StatusBadRequest)
		return
	}

	//pointer to struct for unmarshalling json

	info := &person{
		userInfo{
			Username:  user,
			Password:  pass,
			Firstname: fname,
			Lastname:  lname,
			Blocked:   "true",
		},
	}
	//

	j, err := json.Marshal(info)
	if err != nil {
		http.Error(w, "Can't handle in JSON", 323)
	}

	//Name and create user file
	userFile := user + ".gojson"

	//Directory where data is stored
	db := "user/" + userFile

	//Check if file already exits

	pn, err := ioutil.ReadDir("user/")
	if err != nil {
		log.Fatal(err)
	}
	for _, list := range pn {
		fileNames := string(list.Name())
		s := []string{fileNames}
		fmt.Println(s)
		if helpers.contains(s, userFile) {

		}
	}
	//Creating file to write data into

	_, err = os.Create(db)
	if err != nil {
		http.Error(w, "Cant't create user profile", 500)
	}

	//Write Username Details into user files
	err = ioutil.WriteFile(db, j, 0666)
	if err != nil {
		http.Error(w, "Server can't access user files", 500)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:  "user_id",
			Value: "vokes",
		})
		http.Redirect(w, r, "/home", http.StatusFound)
	}
}
