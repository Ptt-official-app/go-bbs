package main

import (
	// "github.com/PichuChen/go-bbs"
	// "github.com/PichuChen/go-bbs/crypt"
	// "log"
	"net/http"
	"strings"
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

func parseUserPath(path string) (userId string, item string, err error) {
	pathSegment := strings.Split(path, "/")
	// /{{version}}/users/{{user_id}}/{{item}}
	if len(pathSegment) == 4 {
		// /{{version}}/users/{{user_id}}
		return pathSegment[3], "", nil
	}

	return pathSegment[3], pathSegment[4], nil

}
