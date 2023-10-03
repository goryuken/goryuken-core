package controllers

import (
	"fmt"
	"net/http"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Signup")
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Login")
}

func Logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Logout")
}
