package main

import (
	"flag"
	"fmt"

	"github.com/Ptt-official-app/go-bbs"
	_ "github.com/Ptt-official-app/go-bbs/pttbbs"
)

func showArticleList() {
	err := chkIsDir(bbshome)
	if err != nil {
		fmt.Printf("show_article_list: error: %v\n", err)
		return
	}

	db, err := bbs.Open(driverName, "file://"+bbshome)
	if err != nil {
		fmt.Printf("show_article_list: open db: %v\n", err)
		return
	}

	inputArgs := parseArgsToMap(flag.Args())

	// fmt.Println(inputArgs)

	boardID, ok := inputArgs["board"].(string)
	if !ok {
		fmt.Println("please provide board argument")
		return
	}

	records, err := db.ReadBoardArticleRecordsFile(boardID)

	// owner, ok := inputArgs["owner"]
	// if !ok {
	// 	fmt.Println("please provide owner argument")
	// 	return
	// }

	// date, ok := inputArgs["date"]
	// if !ok {
	// 	fmt.Println("please provide date argument")
	// 	return
	// }

	// title, ok := inputArgs["title"]
	// if !ok {
	// 	fmt.Println("please provide title argument")
	// 	return
	// }

	// record, err := db.NewArticleRecord(boardID, owner, date, title)
	// if err != nil {
	// 	fmt.Printf("show_article_list: NewArticleRecord: %v\n", err)
	// 	return
	// }

	for _, r := range records {
		if r.Filename() == "" {
			continue
		}
		fmt.Printf("%s % 12s %s %s\n", r.Filename(), r.Owner(), r.Date(), r.Title())
	}

}
