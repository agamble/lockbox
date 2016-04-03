package main

import (
	"github.com/gorilla/schema"
	"html/template"
	// "log"
	"net/http"
)

func SignUpHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		t, _ := template.ParseFiles("templates/signup.html")
		t.Execute(res, nil)
	} else if req.Method == http.MethodPost {
		err := req.ParseForm()

		if err != nil {
			panic(err)
		}

		tu := &TempUser{}
		dec := schema.NewDecoder()
		dec.Decode(tu, req.PostForm)

		u := NewUser()
		u.Email = tu.Email
		if err = u.Get(); err == nil {
			// user already associated with email address
			panic("Found an existing user while trying to sign up")
		}

		u.SetPassword(tu.Password)
		u.StartTrial()

		err = u.Save()

		if err != nil {
			panic(err)
		}

		t, _ := template.ParseFiles("templates/signup.html")
		t.Execute(res, nil)
	}
}
