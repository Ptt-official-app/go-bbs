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

func main() {
	fmt.Println("server start")

	loadPasswdsFile()

	r := http.NewServeMux()
	r.HandleFunc("/v1/token", routeToken)  // http://127.0.0.1/hello
	r.HandleFunc("/v1/class/", routeClass) // http://127.0.0.1/hello

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
