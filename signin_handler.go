package main

import (
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"html/template"
	// "log"
	"log"
	"net/http"
)

// TODO - change to env variable
var CookieStore *sessions.CookieStore = sessions.NewCookieStore([]byte("something-very-secret"))
var SessionName string = "auth-session"

type TempUser struct {
	Email    string
	Password string
}

func SignInHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		t, _ := template.ParseFiles("templates/signin.html")
		t.Execute(w, nil)
	} else if req.Method == http.MethodPost {
		err := req.ParseForm()
		if err != nil {
			panic(err)
		}

		tu := &TempUser{}
		dec := schema.NewDecoder()
		dec.Decode(tu, req.PostForm)

		log.Printf("%+v", tu)

		u := NewUserFromTemp(tu)

		if err = u.Get(); err != nil {
			// user doesn't exist
			panic(err)
		}

		log.Println("Got user")
		if !u.CorrectPassword(tu.Password) {
			// password wrong
			log.Printf("%+v", u)
			panic("bad password")
		}

		log.Println("correct password")

		session, err := setAuthenticatedCookie(req, u, true)

		if err != nil {
			panic(err)
		}
		log.Println("built session")

		session.Save(req, w)
		t, _ := template.ParseFiles("templates/index.html")
		t.Execute(w, nil)
	}
}

func setAuthenticatedCookie(req *http.Request, u *User, rememberMe bool) (*sessions.Session, error) {
	session, err := CookieStore.Get(req, SessionName)

	var maxAge int = 0

	if rememberMe {
		maxAge = 86400 * 7 * 52
	}

	session.Options = &sessions.Options{
		MaxAge: maxAge,
	}

	if err != nil {
		return nil, err
	}

	session.Values["email"] = u.Email
	return session, nil
}
