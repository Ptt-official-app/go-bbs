package main

import (
	"encoding/json"
	"fmt"
	"github.com/PichuChen/go-bbs"
	"github.com/PichuChen/go-bbs/crypt"
	"log"
	"net/http"
	// "strings"
)

func routeToken(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		postToken(w, r)
		return
	}

}

func postToken(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	username := r.FormValue("username")
	password := r.FormValue("password")

	userec, err := findUserecById(username)
	if err != nil {
		m := map[string]string{
			"error":             "grant_error",
			"error_description": err.Error(),
		}
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Write(b)
		return

	}

	log.Println("found user:", *userec)
	err = verifyPassword(userec, password)
	if err != nil {

		// TODO: add delay, warning, notify user

		m := map[string]string{
			"error":             "grant_error",
			"error_description": err.Error(),
		}
		b, _ := json.MarshalIndent(m, "", "  ")
		w.Write(b)
		return
	}

	// Generate Access Token
	m := map[string]string{
		"access_token": "this-is-mocking-access-token",
		"token_type":   "bearer",
	}

	b, _ := json.MarshalIndent(m, "", "  ")

	w.Write(b)

}

func findUserecById(userid string) (*bbs.Userec, error) {

	for _, it := range userRecs {
		if userid == it.Userid {
			return it, nil
		}
	}
	return nil, fmt.Errorf("user record not found")

}

func verifyPassword(userec *bbs.Userec, password string) error {
	log.Println("password", userec.Passwd)
	res, err := crypt.Fcrypt([]byte(password), []byte(userec.Passwd[:2]))
	str := string(res)
	log.Println("res", str, err)
	if str != userec.Passwd {
		return fmt.Errorf("password incorrect")
	}
	return nil

}
