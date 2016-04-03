package main

import (
	"net/http"

	"github.com/gorilla/context"
	// "github.com/gorilla/sessions"
)

func CreateAuthMw(handler http.Handler) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
	}
}

func Auth(next func(w http.ResponseWriter, req *http.Request)) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := CookieStore.Get(req, SessionName)

		if session.IsNew {
			http.Error(w, "No existing login found", http.StatusUnauthorized)
			return
		}

		email := session.Values["email"]
		emailStr, ok := email.(string)

		if !ok {
			http.Error(w, "Could not decode email", 500)
			return
		}

		user := NewUserFromEmail(emailStr)

		if err = user.Get(); err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		// Set the user value for later access
		context.Set(req, "user", user)
		next(w, req)
	}
}
