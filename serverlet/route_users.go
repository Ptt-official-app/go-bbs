package main

import (
	// "github.com/Ptt-official-app/go-bbs"
	// "github.com/Ptt-official-app/go-bbs/crypt"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func routeUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getUsers(w, r)
		return
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	userID, item, err := parseUserPath(r.URL.Path)

	if item == "information" {
		getUserInformation(w, r, userID)
		return
	} else if item == "favorites" {
		getUserFavorites(w, r, userID)
		return
	}
	// else

	log.Println(userID, item, err)

	w.WriteHeader(http.StatusNotFound)
}

func getUserInformation(w http.ResponseWriter, r *http.Request, userID string) {
	token := getTokenFromRequest(r)
	err := checkTokenPermission(token,
		[]permission{PermissionReadUserInformation},
		map[string]string{
			"user_id": userID,
		})
	if err != nil {
		// TODO: record unauthorized access
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userrec, err := findUserecByID(userID)
	if err != nil {
		// TODO: record error

		w.WriteHeader(http.StatusInternalServerError)
		m := map[string]string{
			"error":             "find_userrec_error",
			"error_description": "get userrec for " + userID + " failed",
		}
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Write(b)
		return
	}

	// TODO: Check Etag or Not-Modified for cache

	dataMap := map[string]interface{}{
		"user_id":              userrec.UserID(),
		"nickname":             userrec.Nickname(),
		"realname":             userrec.RealName(),
		"number_of_login_days": fmt.Sprintf("%d", userrec.NumLoginDays()),
		"number_of_posts":      fmt.Sprintf("%d", userrec.NumPosts()),
		// "number_of_badposts":   fmt.Sprintf("%d", userrec.NumLoginDays),
		"money":           fmt.Sprintf("%d", userrec.Money()),
		"last_login_time": userrec.LastLogin().Format(time.RFC3339),
		"last_login_ipv4": userrec.LastHost(),
		"last_login_ip":   userrec.LastHost(),
		// "last_login_country": fmt.Sprintf("%d", userrec.NumLoginDays),
		"chess_status": map[string]interface{}{},
		"plan":         map[string]interface{}{},
	}

	responseMap := map[string]interface{}{
		"data": dataMap,
	}

	responseByte, _ := json.MarshalIndent(responseMap, "", "  ")

	w.Write(responseByte)
}

func getUserFavorites(w http.ResponseWriter, r *http.Request, userID string) {
	w.WriteHeader(http.StatusNotImplemented)
}

func parseUserPath(path string) (userID string, item string, err error) {
	pathSegment := strings.Split(path, "/")
	// /{{version}}/users/{{user_id}}/{{item}}
	if len(pathSegment) == 4 {
		// /{{version}}/users/{{user_id}}
		return pathSegment[3], "", nil
	}

	return pathSegment[3], pathSegment[4], nil
}
