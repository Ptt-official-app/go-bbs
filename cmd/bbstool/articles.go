package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/Ptt-official-app/go-bbs"
)

func showBoardArticleList() {
	db, err := openDB()
	if err != nil {
		fmt.Println("openDB:", err)
	}
	inputArgs := parseArgsToMap(flag.Args())

	// fmt.Println(inputArgs)

	boardID, ok := inputArgs["board"].(string)
	if !ok {
		fmt.Println("please provide board argument")
		return
	}

	records, err := db.ReadBoardArticleRecordsFile(boardID)

	for _, r := range records {
		if r.Filename() == "" {
			continue
		}
		fmt.Printf("%s % 12s %s %s\n", r.Filename(), r.Owner(), r.Date(), r.Title())
	}
}

func addBoardArticle() {
	db, err := openDB()
	if err != nil {
		fmt.Println("openDB:", err)
	}
	inputArgs := parseArgsToMap(flag.Args())

	boardID, ok := inputArgs["board"].(string)
	if !ok {
		fmt.Println("please provide board argument")
		return
	}

	owner, ok := inputArgs["owner"].(string)
	if !ok {
		fmt.Println("please provide owner argument")
		return
	}

	date, ok := inputArgs["date"].(string)
	if !ok {
		fmt.Println("please provide date argument")
		return
	}

	title, ok := inputArgs["title"].(string)
	if !ok {
		fmt.Println("please provide title argument")
		return
	}

	content, contentOk := inputArgs["content"].(string)
	contentFile, inputFileOk := inputArgs["content_file"].(string)
	if !contentOk && !inputFileOk {
		fmt.Println("please provide content or content_file")
	}

	if !contentOk {
		contentData, err := os.ReadFile(contentFile)
		if err != nil {
			fmt.Println("readfile error:", err)
			return
		}
		content = string(contentData)
	}

	record, err := db.CreateArticleRecord(boardID, owner, date, title)
	if err != nil {
		fmt.Println("show_article_list: NewArticleRecord: ", err)
		return
	}
	err = db.AddArticleRecordFileRecord(boardID, record)
	if err != nil {
		fmt.Println("AddArticleRecordFileRecord error:", err)
		return
	}

	err = db.WriteBoardArticleFile(boardID, record.Filename(), bbs.Utf8ToBig5(content))
	if err != nil {
		fmt.Println("WriteBoardArticleFile error: ", err)
		return
	}
	fmt.Printf("new filename: %s\n", record.Filename())
	return

}

func showBoardArticle() {
	db, err := openDB()
	if err != nil {
		fmt.Println("openDB:", err)
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
