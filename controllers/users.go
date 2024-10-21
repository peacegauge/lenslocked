package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {

	// you are getting the names you assigned each input in you form
	fmt.Fprint(w, "Email: ", r.FormValue("email"), "\n")
	fmt.Fprint(w, "Password: ", r.FormValue("password"))
}
