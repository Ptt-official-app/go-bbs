package main

import (
	"flag"
	"fmt"

	"github.com/Ptt-official-app/go-bbs"
)

func showBoardArticleList() {
	db, err := openDB()
	if err != nil {
		fmt.Printf("openDB: %w", err)
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

func showBoardArticle() {
	db, err := openDB()
	if err != nil {
		fmt.Printf("openDB: %w", err)
	}
	inputArgs := parseArgsToMap(flag.Args())

	boardID, ok := inputArgs["board"].(string)
	if !ok {
		fmt.Println("please provide board argument")
		return
	}

	filename, ok := inputArgs["filename"].(string)
	if !ok {
		fmt.Println("please provide filename argument")
		return
	}

	data, err := db.ReadBoardArticleFile(boardID, filename)
	if err != nil {
		fmt.Println("ReadrBoardArticleFile error: %w", err)
		return
	}
	fmt.Printf("filecontent: \n%s\n", bbs.Big5ToUtf8(data))
	return

}
