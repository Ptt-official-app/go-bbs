package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Ptt-official-app/go-bbs"
	_ "github.com/Ptt-official-app/go-bbs/pttbbs"
)

var driverName string
var bbshome string

func init() {
	flag.StringVar(&driverName, "driver", "pttbbs", "bbs driver name, eg: pttbbs, maple, formosabbs")
	flag.StringVar(&bbshome, "bbshome", "/home/bbs", "bbshome path, default is /home/bbs")
}

func main() {
	flag.Parse()

	fmt.Printf("driver name: %v\n", driverName)
	fmt.Printf("bbshome: %v\n", bbshome)
	fmt.Printf("args: %v\n", flag.Args())

	switch flag.Arg(0) {
	case "show_board_list":
		showboardlist()
		return
	case "add_board":
		addboard()
		return
	case "show_user_list":
		showuserlist()
		return
	case "show_user_article_list":
		// Example: go run ./  --bbshome=../../home/bbs show_user_article_list --user_id pichu
		showuserarticlelist()
		return
	case "show_user_comment_list":
		// Example: go run ./  --bbshome=../../home/bbs show_user_comment_list --user_id pichu --board t
		showusercommentlist()
		return

	case "show_board_article_list":
		// Example: go run ./  --bbshome=../../home/bbs show_board_article_list --board test
		showBoardArticleList()
		return

	case "add_board_article":
		// Example: go run ./  --bbshome=../../home/bbs add_board_article --board test
		addBoardArticle()
		return

	case "show_board_article":
		// Example: go run ./  --bbshome=../../home/bbs show_board_article --board test --filename M.1621323571.A.4C6
		showBoardArticle()
		return
	default:
		fmt.Printf("unknown subcommand: %s\n", flag.Arg(0))
		return
	}

}

func chkIsDir(dir string) error {
	file, err := os.Open(dir)
	if err != nil {
		return err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return err
	}

	if stat.IsDir() {
		return nil
	}
	return fmt.Errorf("%v is not dir", dir)

}

func openDB() (db *bbs.DB, err error) {
	err = chkIsDir(bbshome)
	if err != nil {
		// fmt.Printf("showBoardArticleList: error: %v\n", err)
		return
	}

	db, err = bbs.Open(driverName, "file://"+bbshome)
	if err != nil {
		// fmt.Printf("showBoardArticleList: open db: %v\n", err)
		return
	}
	return
}

func parseArgsToMap(flagArgs []string) map[string]interface{} {
	newBoardArgs := map[string]interface{}{}
	index := 1
	for {
		if index >= len(flagArgs) {
			// fmt.Println("parsed")
			break
		}

		if strings.HasPrefix(flagArgs[index], "--") {
			if index+1 >= len(flagArgs) {
				fmt.Println("flag with no args", flagArgs[index])
				break
			}
			newBoardArgs[flagArgs[index][2:]] = flagArgs[index+1]
			index = index + 2
			continue
		}
		index = index + 1

	}
	return newBoardArgs
}

func showboardlist() {
	err := chkIsDir(bbshome)
	if err != nil {
		fmt.Printf("showboardlist: error: %v\n", err)
		return
	}

	bbsDB, err := bbs.Open(driverName, "file://"+bbshome)
	if err != nil {
		fmt.Printf("showboardlist: open db: %v\n", err)
		return
	}

	records, err := bbsDB.ReadBoardRecords()
	if err != nil {
		fmt.Printf("showboardlist: ReadBoardRecords: %v\n", err)
		return
	}
	for _, r := range records {
		fmt.Println("title:", r.Title())
	}

}

func addboard() {
	err := chkIsDir(bbshome)
	if err != nil {
		fmt.Printf("addboard: error: %v\n", err)
		return
	}

	bbsDB, err := bbs.Open(driverName, "file://"+bbshome)
	if err != nil {
		fmt.Printf("addboard: open db: %v\n", err)
		return
	}

	newBoardArgs := parseArgsToMap(flag.Args())

	fmt.Printf("addboard args: %v\n", newBoardArgs)

	newRecord, err := bbsDB.NewBoardRecord(newBoardArgs)
	if err != nil {
		fmt.Printf("addboard: NewBoardRecord: %v\n", err)
		return
	}

	err = bbsDB.AddBoardRecord(newRecord)
	if err != nil {
		fmt.Printf("addboard: AddBoardRecordFileRecord: %v\n", err)
		return
	}

}

func showuserlist() {
	err := chkIsDir(bbshome)
	if err != nil {
		fmt.Printf("showuserlist: error: %v\n", err)
		return
	}

	bbsDB, err := bbs.Open(driverName, "file://"+bbshome)
	if err != nil {
		fmt.Printf("showuserlist: open db: %v\n", err)
		return
	}

	records, err := bbsDB.ReadUserRecords()
	if err != nil {
		fmt.Printf("showuserlist: ReadUserRecords: %v\n", err)
		return
	}

	for _, r := range records {
		if r.UserID() == "" {
			continue
		}
		fmt.Println("user id:", r.UserID())
	}

}

func showuserarticlelist() {
	err := chkIsDir(bbshome)
	if err != nil {
		fmt.Printf("showuserarticlelist: error: %v\n", err)
		return
	}

	bbsDB, err := bbs.Open(driverName, "file://"+bbshome)
	if err != nil {
		fmt.Printf("showuserarticlelist: open db: %v\n", err)
		return
	}

	args := parseArgsToMap(flag.Args())

	records, err := bbsDB.GetUserArticleRecordFile(args["user_id"].(string))
	if err != nil {
		fmt.Printf("showuserarticlelist: GetUserArticleRecordFile: %v\n", err)
		return
	}

	for _, r := range records {
		// if r.UserID() == "" {
		// 	continue
		// }
		fmt.Println("titlex:", r.Title())
	}

}

func showusercommentlist() {
	err := chkIsDir(bbshome)
	if err != nil {
		fmt.Printf("showusercommentlist: error: %v\n", err)
		return
	}

	bbsDB, err := bbs.Open(driverName, "file://"+bbshome)
	if err != nil {
		fmt.Printf("showusercommentlist: open db: %v\n", err)
		return
	}

	args := parseArgsToMap(flag.Args())
	userID := args["user_id"].(string)

	records, err := bbsDB.GetUserCommentRecordFile(userID)
	if err != nil {
		fmt.Printf("showusercommentlist: GetUserCommentRecordFile: %v\n", err)
		return
	}

	for _, r := range records {
		fmt.Println("user comment record:", r.String())
	}
}
