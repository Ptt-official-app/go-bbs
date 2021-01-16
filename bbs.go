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

type FavoriteType int

const (
	FavoriteTypeBoard  FavoriteType = iota // 0
	FavoriteTypeFolder                     // 1
	FavoriteTypeLine                       // 2

)

type FavoriteRecord interface {
	Title() string
	Type() FavoriteType
	BoardId() string

	// Records is FavoriteTypeFolder only.
	Records() []FavoriteRecord
}

type BoardRecord interface {
	BoardId() string

	Title() string

	IsClass() bool

	BM() []string
}

type ArticleRecord interface {
	Filename() string
	Modified() time.Time
	Recommend() int
	Date() string
	Title() string
	Money() int
	Owner() string
}

// DB is whole bbs filesystem, including where file store,
// how to connect to local cache ( system V shared memory or etc.)
// how to parse or store it's data to bianry
type DB struct {
	connector Connector
}

// Driver should implement Connector interface
type Connector interface {
	// Open provides the driver parameter settings, such as BBSHome parameter and SHM parameters.
	Open(dataSourceName string) error
	// GetUserRecordsPath should return user records file path, eg: BBSHome/.PASSWDS
	GetUserRecordsPath() (string, error)
	// ReadUserRecordsFile should return UserRecord list in the file called name
	ReadUserRecordsFile(name string) ([]UserRecord, error)
	// GetUserFavoriteRecordsPath should return the user favorite records file path
	// for specific user, eg: BBSHOME/home/{{u}}/{{userId}}/.fav
	GetUserFavoriteRecordsPath(userId string) (string, error)
	// ReadUserFavoriteRecordsFile should return FavoriteRecord list in the file called name
	ReadUserFavoriteRecordsFile(name string) ([]FavoriteRecord, error)
	// GetBoardRecordsPath should return the board headers file path, eg: BBSHome/.BRD
	GetBoardRecordsPath() (string, error)
	// ReadBoardRecordsFile shoule return BoardRecord list in file, name is the file name
	ReadBoardRecordsFile(name string) ([]BoardRecord, error)
	// GetBoardArticleRecordsPath should return the article records file path, boardId is the board id,
	// eg: BBSHome/boards/{{b}}/{{boardId}}/.DIR
	GetBoardArticleRecordsPath(boardId string) (string, error)
	// GetBoardArticleRecordsPath should return the treasure records file path, boardId is the board id,
	// eg: BBSHome/man/boards/{{b}}/{{boardId}}/{{treasureId}}/.DIR
	GetBoardTreasureRecordsPath(boardId string, treasureId []string) (string, error)
	// ReadArticleRecordsFile returns ArticleRecord list in file, name is the file name
	ReadArticleRecordsFile(name string) ([]ArticleRecord, error)
	// GetBoardArticleFilePath return file path for specific boardId and filename
	GetBoardArticleFilePath(boardId string, filename string) (string, error)
	// GetBoardTreasureFilePath return file path for specific boardId, treasureId and filename
	GetBoardTreasureFilePath(boardId string, treasureId []string, name string) (string, error)
	// ReadBoardArticleFile should returns raw file of specific file name
	ReadBoardArticleFile(name string) ([]byte, error)
}

var drivers = make(map[string]Connector)

func Register(drivername string, connector Connector) {
	// TODO: Mutex
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

// ReadUserFavoriteRecords returns the FavoriteRecord for specific userId
func (db *DB) ReadUserFavoriteRecords(userId string) ([]FavoriteRecord, error) {

	path, err := db.connector.GetUserFavoriteRecordsPath(userId)
	if err != nil {
		log.Println("bbs: get user favorite records path error:", err)
		return nil, err
	}
	log.Println("path:", path)

	recs, err := db.connector.ReadUserFavoriteRecordsFile(path)
	if err != nil {
		log.Println("bbs: read user favorite records error:", err)
		return nil, err
	}
	return recs, nil

	return nil, err

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

func (db *DB) ReadBoardArticleRecordsFile(boardId string) ([]ArticleRecord, error) {

	path, err := db.connector.GetBoardArticleRecordsPath(boardId)
	if err != nil {
		log.Println("bbs: open file error:", err)
		return nil, err
	}
	log.Println("path:", path)

	recs, err := db.connector.ReadArticleRecordsFile(path)
	if err != nil {
		log.Println("bbs: get user rec error:", err)
		return nil, err
	}
	return recs, nil

}

func (db *DB) ReadBoardTreasureRecordsFile(boardId string, treasureId []string) ([]ArticleRecord, error) {

	path, err := db.connector.GetBoardTreasureRecordsPath(boardId, treasureId)
	if err != nil {
		log.Println("bbs: open file error:", err)
		return nil, err
	}
	log.Println("path:", path)

	recs, err := db.connector.ReadArticleRecordsFile(path)
	if err != nil {
		log.Println("bbs: get user rec error:", err)
		return nil, err
	}
	return recs, nil
}

func (db *DB) ReadBoardArticleFile(boardId string, filename string) ([]byte, error) {

	path, err := db.connector.GetBoardArticleFilePath(boardId, filename)
	if err != nil {
		log.Println("bbs: open file error:", err)
		return nil, err
	}
	log.Println("path:", path)

	recs, err := db.connector.ReadBoardArticleFile(path)
	if err != nil {
		log.Println("bbs: get user rec error:", err)
		return nil, err
	}
	return recs, nil
}

func (db *DB) ReadBoardTreasureFile(boardId string, treasuresId []string, filename string) ([]byte, error) {

	path, err := db.connector.GetBoardTreasureFilePath(boardId, treasuresId, filename)
	if err != nil {
		log.Println("bbs: open file error:", err)
		return nil, err
	}
	log.Println("path:", path)

	recs, err := db.connector.ReadBoardArticleFile(path)
	if err != nil {
		log.Println("bbs: get user rec error:", err)
		return nil, err
	}
	return recs, nil
}
