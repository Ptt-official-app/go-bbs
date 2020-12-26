package main

import (
	// "github.com/PichuChen/go-bbs"
	// "github.com/PichuChen/go-bbs/crypt"
	// "log"
	"net/http"
	// "strings"
)

func routeClasses(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getClasses(w, r)
		return
	}

}

func getClasses(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotImplemented)

}
