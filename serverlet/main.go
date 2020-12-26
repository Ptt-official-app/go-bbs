package main

import (
	"encoding/json"
	"fmt"
	"github.com/PichuChen/go-bbs"
	"log"
	"net/http"
	"strings"
)

var userRecs []*bbs.Userec
var boardHeader []*bbs.BoardHeader

func main() {
	fmt.Println("server start")

	loadPasswdsFile()
	loadBoardFile()

	r := http.NewServeMux()
	r.HandleFunc("/v1/token", routeToken)
	r.HandleFunc("/v1/boards", routeBoards)
	r.HandleFunc("/v1/classes/", routeClasses)
	r.HandleFunc("/v1/users/", routeUsers)

	http.ListenAndServe(":8080", r)
}

func loadPasswdsFile() {
	// TODO: read config form config file

	path, err := bbs.GetPasswdsPath("../home/bbs")
	if err != nil {
		log.Println("open file error:", err)
		return
	}
	log.Println("path:", path)

	userRecs, err = bbs.OpenUserecFile(path)
	if err != nil {
		log.Println("get user rec error:", err)
		return
	}
	log.Println(userRecs)
}

func loadBoardFile() {
	// TODO: read config form config file

	path, err := bbs.GetBoardPath("../home/bbs")
	if err != nil {
		log.Println("open file error:", err)
		return
	}
	log.Println("path:", path)

	boardHeader, err = bbs.OpenBoardHeaderFile(path)
	if err != nil {
		log.Println("get board header error:", err)
		return
	}
	log.Println(userRecs)
}

func routeClass(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getClass(w, r)
		return
	}

}

func getClass(w http.ResponseWriter, r *http.Request) {

	seg := strings.Split(r.URL.Path, "/")

	classId := "0"
	if len(seg) > 2 {
		classId = seg[3]
	}

	log.Println("user get class:", classId)

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
