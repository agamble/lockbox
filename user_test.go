package main

import "testing"

// func TestLoad(t *testing.T) {
// 	u := NewUser()
//
// 	u.Email = "alexander.gamble@outlook.com"
// 	err := u.Save()
//
// 	if err != nil {
// 		panic(err)
// 	}
//
// }

func TestPasswordStore(t *testing.T) {
	u := NewUser()
	u.Password = "hermione"
	u.HashPassword()
	if !u.CorrectPassword() {
		panic("Bad password")
	}
}
