package main

import (
	"fmt"

	"github.com/Ptt-official-app/go-bbs/pttbbs"
)

func main() {
	c, err := pttbbs.NewCache("file://../../dump.shm", &pttbbs.MemoryMappingSetting{
		MaxUsers: 50,
		IDLen:    12,
	})
	if err != nil {
		fmt.Println("new cache err:", err)
		return
	}
	fmt.Println("version", c.Version())
	fmt.Println("userid, 0:", c.UserID(0))
	fmt.Println("money, 0:", c.Money(0))
}
