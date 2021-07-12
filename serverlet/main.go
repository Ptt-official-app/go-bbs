package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Ptt-official-app/go-bbs"
	_ "github.com/Ptt-official-app/go-bbs/pttbbs"
)

var (
	userRecs    []bbs.UserRecord
	boardHeader []bbs.BoardRecord
)

var bbsDB *bbs.DB

func main() {
	fmt.Println("server start")
	var err error
	bbsDB, err = bbs.Open("pttbbs", "file://../home/bbs")
	if err != nil {
		log.Printf("open bbs error: %v", err)
		return
	}

	loadPasswdsFile()
	loadBoardFile()

	r := http.NewServeMux()
	r.HandleFunc("/v1/token", routeToken)
	r.HandleFunc("/v1/boards", routeBoards)
	r.HandleFunc("/v1/classes/", routeClasses)
	r.HandleFunc("/v1/users/", routeUsers)

	http.ListenAndServe(":8083", r)
}

func loadPasswdsFile() {
	var err error
	userRecs, err = bbsDB.ReadUserRecords()
	if err != nil {
		log.Println("get user rec error:", err)
		return
	}
	log.Println(userRecs)
}

func loadBoardFile() {
	var err error
	boardHeader, err = bbsDB.ReadBoardRecords()
	if err != nil {
		log.Println("get board header error:", err)
		return
	}
	log.Println(boardHeader)
}

func routeClass(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getClass(w, r)
		return
	}
}

func getClass(w http.ResponseWriter, r *http.Request) {
	seg := strings.Split(r.URL.Path, "/")

	classID := "0"
	if len(seg) > 2 {
		classID = seg[3]
	}

	log.Println("user get class:", classID)

	list := []interface{}{}

	c := map[string]interface{}{
		"id":             1,
		"type":           "class",
		"title":          "title",
		"number_of_user": 3,
		"moderators": []string{
			"SYSOP",
			"pichu",
		},
	}
	list = append(list, c)

	m := map[string]interface{}{
		"data": list,
	}
	b, _ := json.MarshalIndent(m, "", "  ")

	w.Write(b)
}
