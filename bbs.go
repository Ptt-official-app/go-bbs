package bbs

import (
	"fmt"
	"log"
	"time"
)

// UserRecord mapping to `userec` in most system, it records uesr's
// basical data
type UserRecord interface {
	// UserId return user's identification string, and it is userid in
	// mostly bbs system
	UserId() string
	// HashedPassword return user hashed password, it only for debug,
	// If you want to check is user password correct, please use
	// VerifyPassword insteaded.
	HashedPassword() string
	// VerifyPassword will check user's password is OK. it will return null
	// when OK and error when there are something wrong
	VerifyPassword(password string) error
	// Nickname return a string for user's nickname, this string may change
	// depend on user's mood, return empty string if this bbs system do not support
	Nickname() string
	// RealName return a string for user's real name, this string may not be changed
	// return empty string if this bbs system do not support
	RealName() string
	// NumLoginDays return how many days this have been login since account created.
	NumLoginDays() int
	// NumPosts return how many posts this user has posted.
	NumPosts() int
	// Money return the money this user have.
	Money() int
	// LastLogin return last login time of user
	LastLogin() time.Time
	// LastHost return last login host of user, it is IPv4 address usually, but it
	// could be domain name or IPv6 address.
	LastHost() string
}

type BoardRecord interface {
	BoardId() string

	Title() string

	IsClass() bool

	BM() []string
}

type ArticleRecord interface{}

// DB is whole bbs filesystem, including where file store,
// how to connect to local cache ( system V shared memory or etc.)
// how to parse or store it's data to bianry
type DB struct {
	connector Connector
}

type Connector interface {
	Open(dataSourceName string) error

	GetUserRecordsPath() (string, error)
	ReadUserRecordsFile(name string) ([]UserRecord, error)

	GetBoardRecordsPath() (string, error)
	ReadBoardRecordsFile(name string) ([]BoardRecord, error)

	GetBoardArticleRecordsPath(boardId string) (string, error)
	ReadBoardArticleRecordsFile(boardId string) ([]ArticleRecord, error)

	GetBoardTreasureRecordsPath(boardId string, treasuresId []string) (string, error)
	ReadBoardTreasureRecordsFile(boardId string, treasuresId []string) ([]ArticleRecord, error)

	ReadBoardArticleFile(boardId, filename string) ([]byte, error)
	ReadBoardTreasureFile(boardId string, treasuresId []string, filename string) ([]byte, error)
}

var drivers = make(map[string]Connector)

func Register(drivername string, connector Connector) {
	drivers[drivername] = connector
}

// Open opan a
func Open(drivername string, dataSourceName string) (*DB, error) {

	c, ok := drivers[drivername]
	if !ok {
		return nil, fmt.Errorf("bbs: drivername: %v not found", drivername)
	}

	err := c.Open(dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("bbs: drivername: %v open error: %v", drivername, err)
	}

	return &DB{
		connector: c,
	}, nil
}

// ReadUserRecords returns the UserRecords
func (db *DB) ReadUserRecords() ([]UserRecord, error) {

	path, err := db.connector.GetUserRecordsPath()
	if err != nil {
		log.Println("bbs: open file error:", err)
		return nil, err
	}
	log.Println("path:", path)

	userRecs, err := db.connector.ReadUserRecordsFile(path)
	if err != nil {
		log.Println("bbs: get user rec error:", err)
		return nil, err
	}
	return userRecs, nil
}

// ReadBoardRecords returns the UserRecords
func (db *DB) ReadBoardRecords() ([]BoardRecord, error) {

	path, err := db.connector.GetBoardRecordsPath()
	if err != nil {
		log.Println("bbs: open file error:", err)
		return nil, err
	}
	log.Println("path:", path)

	recs, err := db.connector.ReadBoardRecordsFile(path)
	if err != nil {
		log.Println("bbs: get user rec error:", err)
		return nil, err
	}
	return recs, nil
}
