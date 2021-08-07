package main

import (
	"encoding/json"
	"net/http"
)

func routeBoards(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getBoards(w, r)
		return
	}

}

func getBoards(w http.ResponseWriter, r *http.Request) {

	// TODO: Check JWT

	// TODO: Get user Level

	// TODO: Show Board by user level

	dataList := []interface{}{}
	for _, b := range boardHeader {
		dataList = append(dataList, b)
	}

	responseMap := map[string]interface{}{
		"data": dataList,
	}

	b, _ := json.MarshalIndent(responseMap, "", "  ")
	w.Write(b)

}
