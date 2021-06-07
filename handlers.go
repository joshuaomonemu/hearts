package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"main.go/helpers"
)

//Map that stores values read from GOJSON
var passer map[string]interface{}

//Handles Sign-in Request
func signinAction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request", 400)
	}
	user := r.FormValue("username")
	password := r.FormValue("password")

	userFile := "user/" + user + ".gojson"
	//Checking if file exits before reading;

	//Reading the user directory
	fi, _ := ioutil.ReadDir("user/")
	//Looping through to check

	//Variables to use while looping
	var f os.FileInfo
	var sf string
	for _, f = range fi {
		sf = f.Name()
		fmt.Println(sf)
	}
	rep := strings.Contains(userFile, sf)
	fmt.Println(rep)
	if rep == true {
		http.Error(w, "Cannot find userfile", http.StatusMovedPermanently)
		return
	}

	reply := helpers.Reader(userFile)

	//Converting GOJSON values and storing in a MAP
	helpers.Unmarshal(reply, &passer)

	//Check if username and password are valid
	userField := passer["person"].(map[string]interface{})["username"]
	passField := passer["person"].(map[string]interface{})["password"]
	block := passer["person"].(map[string]interface{})["blocked"]

	//Converting needed values to string

	acctBlock := fmt.Sprintf("%v", block)

	//What happens if account is blocked
	if acctBlock == "true" {
		http.Redirect(w, r, "/error", http.StatusFound)
		return
	}
	//Conditional statements to ensure struct values are the same
	if passField == password && userField == user {
		c := &http.Cookie{
			Name:  "user_id",
			Value: user,
		}
		http.SetCookie(w, c)
		http.Redirect(w, r, "/main", http.StatusFound)
	} else {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, `<h1>Incorrect Password Or Username</h1>`)
	}
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
		if helpers.Contains(s, userFile) {
			http.Error(w, "This name already exits", 400)
			return
		}
	}

	//Creating file to write data into

	_, err = os.Create(db)
	if err != nil {
		http.Error(w, "Can't create user profile", 500)
	}

	//Write Username Details into user files
	err = ioutil.WriteFile(db, j, 0666)
	if err != nil {
		http.Error(w, "Server can't access user files", 500)
	} else {
		http.SetCookie(w, &http.Cookie{
			Name:  "user_id",
			Value: user,
		})
		http.Redirect(w, r, "/main", http.StatusFound)
	}
}

func renderUser(w http.ResponseWriter, r *http.Request) {
	user, err := r.Cookie("user_id")
	if err != nil {
		http.Redirect(w, r, "/signin", http.StatusMovedPermanently)
		return
	}
	//Cookie Value
	cv := user.Value

	userFile := "user/" + cv + ".gojson"

	//Read User File

	bs := helpers.Reader(userFile)

	//Convert data that has been read in GOJSON to a map-interface
	helpers.Unmarshal(bs, &passer)
	fname := passer["person"].(map[string]interface{})["firstname"]
	lname := passer["person"].(map[string]interface{})["lastname"]
	uname := passer["person"].(map[string]interface{})["username"]

	//Converting values to string to store in a struct
	firstname := fmt.Sprintf("%v", fname)
	lastname := fmt.Sprintf("%v", lname)
	username := fmt.Sprintf("%v", uname)

	//Storing Read values in struct to pass into template
	user_details := person{
		userInfo{
			Firstname: firstname,
			Lastname:  lastname,
			Username:  username,
		},
	}
	//fmt.Println(firstname)
	tpl, err := template.ParseFiles("templates/main.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.Execute(w, user_details.Person)

}
func logOut(w http.ResponseWriter, r *http.Request) {
	c := http.Cookie{
		Name:   "user_id",
		MaxAge: -1}
	http.SetCookie(w, &c)

	http.Redirect(w, r, "/", http.StatusFound)
}

//404 Page
func lostPage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/pages/404-page.gohtml")
	t.Execute(w, nil)
}

//End 404 page
