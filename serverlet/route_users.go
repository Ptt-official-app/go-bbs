package main

import (
	// "github.com/PichuChen/go-bbs"
	// "github.com/PichuChen/go-bbs/crypt"
	// "log"
	"net/http"
	// "strings"
)

func routeUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getUsers(w, r)
		return
	}

}

func getUsers(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotImplemented)
}
