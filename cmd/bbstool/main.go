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
	case "showboardlist":
		showboardlist()
		return
	case "addboard":
		addboard()
		return
	case "showuserlist":
		showuserlist()
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
		if r.UserId() == "" {
			continue
		}
		fmt.Println("user id:", r.UserId())
	}

}
