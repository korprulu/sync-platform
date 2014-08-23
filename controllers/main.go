package controllers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
}

func SignIn(w http.ResponseWriter, r *http.Request) {
}
