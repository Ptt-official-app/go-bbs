package bbs

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
)

// UserRecord mapping to `userec` in most system, it records uesr's
// basical data
type UserRecord interface {
	// UserID return user's identification string, and it is userid in
	// mostly bbs system
	UserID() string
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
	// UserFlag return user setting.
	// uint32, see https://github.com/ptt/pttbbs/blob/master/include/uflags.h
	UserFlag() uint32
}

// BadPostUserRecord return UserRecord interface which support NumBadPosts
type BadPostUserRecord interface {
	// NumBadPosts return how many bad post this use have
	NumBadPosts() int
}

// LastCountryUserRecord return UserRecord interface which support LastCountry
type LastCountryUserRecord interface {
	// LastLoginCountry will return the country with this user's last login IP
	LastLoginCountry() string
}

// MailboxUserRecord return UserRecord interface which support MailboxDescription
type MailboxUserRecord interface {
	// MailboxDescription will return the mailbox description with this user
	MailboxDescription() string
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
	BoardID() string

	// Records is FavoriteTypeFolder only.
	Records() []FavoriteRecord
}

type BoardRecord interface {
	BoardID() string

	Title() string

	IsClass() bool
	// ClassID should return the class id to which this board/class belongs.
	ClassID() string

	BM() []string
}

type BoardRecordSettings interface {
	IsHide() bool
	IsPostMask() bool
	IsAnonymous() bool
	IsDefaultAnonymous() bool
	IsNoCredit() bool
	IsVoteBoard() bool
	IsWarnEL() bool
	IsTop() bool
	IsNoRecommend() bool
	IsAngelAnonymous() bool
	IsBMCount() bool
	IsNoBoo() bool
	IsRestrictedPost() bool
	IsGuestPost() bool
	IsCooldown() bool
	IsCPLog() bool
	IsNoFastRecommend() bool
	IsIPLogRecommend() bool
	IsOver18() bool
	IsNoReply() bool
	IsAlignedComment() bool
	IsNoSelfDeletePost() bool
	IsBMMaskContent() bool
}

type BoardRecordInfo interface {
	GetPostLimitPosts() uint8
	GetPostLimitLogins() uint8
	GetPostLimitBadPost() uint8
}

type ArticleRecord interface {
	Filename() string
	Modified() time.Time
	SetModified(newModified time.Time)
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
	// for specific user, eg: BBSHOME/home/{{u}}/{{userID}}/.fav
	GetUserFavoriteRecordsPath(userID string) (string, error)
	// ReadUserFavoriteRecordsFile should return FavoriteRecord list in the file called name
	ReadUserFavoriteRecordsFile(name string) ([]FavoriteRecord, error)
	// GetBoardRecordsPath should return the board headers file path, eg: BBSHome/.BRD
	GetBoardRecordsPath() (string, error)
	// ReadBoardRecordsFile shoule return BoardRecord list in file, name is the file name
	ReadBoardRecordsFile(name string) ([]BoardRecord, error)
	// GetBoardArticleRecordsPath should return the article records file path, boardID is the board id,
	// eg: BBSHome/boards/{{b}}/{{boardID}}/.DIR
	GetBoardArticleRecordsPath(boardID string) (string, error)
	// GetBoardArticleRecordsPath should return the treasure records file path, boardID is the board id,
	// eg: BBSHome/man/boards/{{b}}/{{boardID}}/{{treasureID}}/.DIR
	GetBoardTreasureRecordsPath(boardID string, treasureID []string) (string, error)
	// ReadArticleRecordsFile returns ArticleRecord list in file, name is the file name
	ReadArticleRecordsFile(name string) ([]ArticleRecord, error)
	// GetBoardArticleFilePath return file path for specific boardID and filename
	GetBoardArticleFilePath(boardID string, filename string) (string, error)
	// GetBoardTreasureFilePath return file path for specific boardID, treasureID and filename
	GetBoardTreasureFilePath(boardID string, treasureID []string, name string) (string, error)
	// ReadBoardArticleFile should returns raw file of specific file name
	ReadBoardArticleFile(name string) ([]byte, error)
}

// Driver which implement WriteBoardConnector supports modify board record file.
type WriteBoardConnector interface {

	// NewBoardRecord return BoardRecord object in this driver with arguments
	NewBoardRecord(args map[string]interface{}) (BoardRecord, error)

	// AddBoardRecordFileRecord given record file name and new record, should append
	// file record in that file.
	AddBoardRecordFileRecord(name string, brd BoardRecord) error

	// UpdateBoardRecordFileRecord update boardRecord brd on index in record file,
	// index is start with 0
	UpdateBoardRecordFileRecord(name string, index uint, brd BoardRecord) error

	// ReadBoardRecordFileRecord return boardRecord brd on index in record file.
	ReadBoardRecordFileRecord(name string, index uint) (BoardRecord, error)

	// RemoveBoardRecordFileRecord remove boardRecord brd on index in record file.
	RemoveBoardRecordFileRecord(name string, index uint) error
}

// WriteArticleConnector is a connector for writing a article
type WriteArticleConnector interface {

	// CreateBoardArticleFilename returns available filename for board with boardID
	CreateBoardArticleFilename(boardID string) (filename string, err error)

	// NewArticleRecord return ArticleRecord object in this driver with arguments
	NewArticleRecord(filename, owner, date, title string) (ArticleRecord, error)

	// AddArticleRecordFileRecord given record file name and new record, should append
	// file record in that file.
	AddArticleRecordFileRecord(name string, article ArticleRecord) error

	// UpdateArticleRecordFileRecord will write article in position index of name file
	// position is start with 0.
	UpdateArticleRecordFileRecord(name string, index uint, article ArticleRecord) error

	// WriteBoardArticleFile will turncate name file and write content into that file.
	WriteBoardArticleFile(name string, content []byte) error

	// AppendNewLine append content into file
	AppendBoardArticleFile(name string, content []byte) error
}

// UserArticleConnector is a connector for bbs who support cached user article records
type UserArticleConnector interface {

	// GetUserArticleRecordsPath should return the file path which user article record stores.
	GetUserArticleRecordsPath(userID string) (string, error)

	// ReadUserArticleRecordFile should return the article record in file.
	ReadUserArticleRecordFile(name string) ([]UserArticleRecord, error)

	// WriteUserArticleRecordFile write user article records into file.
	WriteUserArticleRecordFile(name string, records []UserArticleRecord) error

	// AppendUserArticleRecordFile append user article records into file.
	AppendUserArticleRecordFile(name string, record UserArticleRecord) error
}

// UserCommentConnector is a connector for bbs to access the cached user
// comment records.
type UserCommentConnector interface {

	// GetUserCommentRecordsPath should return the file path where storing the
	//  user comment records.
	GetUserCommentRecordsPath(userID string) (string, error)

	// ReadUserCommentRecordFile should return the use comment records from the
	//  file.
	ReadUserCommentRecordFile(name string) ([]UserCommentRecord, error)
}

// UserDraftConnector is a connector for bbs which supports modify user
// draft file.
type UserDraftConnector interface {

	// GetUserDraftPath should return the user's draft file path
	// eg: BBSHome/home/{{u}}/{{userID}}/buf.{{draftID}}
	GetUserDraftPath(userID, draftID string) (string, error)

	// ReadUserDraft return the user draft from the named file.
	ReadUserDraft(name string) ([]byte, error)

	// DeleteUserDraft should remove the named file.
	DeleteUserDraft(name string) error

	// WriteUserDraft should replace user draft from named file and user draft data
	WriteUserDraft(name string, draft []byte) error
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

// ReadUserFavoriteRecords returns the FavoriteRecord for specific userID
func (db *DB) ReadUserFavoriteRecords(userID string) ([]FavoriteRecord, error) {

	path, err := db.connector.GetUserFavoriteRecordsPath(userID)
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

func (db *DB) ReadBoardArticleRecordsFile(boardID string) ([]ArticleRecord, error) {

	path, err := db.connector.GetBoardArticleRecordsPath(boardID)
	if err != nil {
		log.Println("bbs: open file error:", err)
		return nil, err
	}
	log.Println("path:", path)

	recs, err := db.connector.ReadArticleRecordsFile(path)
	if err != nil {
		if strings.Contains(err.Error(), "no such file or directory") {
			return []ArticleRecord{}, nil
		}
		log.Println("bbs: ReadArticleRecordsFile error:", err)
		return nil, err
	}
	return recs, nil

}

func (db *DB) ReadBoardTreasureRecordsFile(boardID string, treasureID []string) ([]ArticleRecord, error) {

	path, err := db.connector.GetBoardTreasureRecordsPath(boardID, treasureID)
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

func (db *DB) ReadBoardArticleFile(boardID string, filename string) ([]byte, error) {

	path, err := db.connector.GetBoardArticleFilePath(boardID, filename)
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

func (db *DB) ReadBoardTreasureFile(boardID string, treasuresID []string, filename string) ([]byte, error) {

	path, err := db.connector.GetBoardTreasureFilePath(boardID, treasuresID, filename)
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

func (db *DB) NewBoardRecord(args map[string]interface{}) (BoardRecord, error) {
	return db.connector.(WriteBoardConnector).NewBoardRecord(args)
}

func (db *DB) AddBoardRecord(brd BoardRecord) error {

	path, err := db.connector.GetBoardRecordsPath()
	if err != nil {
		log.Println("bbs: open file error:", err)
		return err
	}
	log.Println("path:", path)

	err = db.connector.(WriteBoardConnector).AddBoardRecordFileRecord(path, brd)
	if err != nil {
		log.Println("bbs: AddBoardRecordFileRecord error:", err)
		return err
	}
	return nil
}

// UpdateBoardRecordFileRecord update boardRecord brd on index in record file,
// index is start with 0
func (db *DB) UpdateBoardRecord(index uint, brd *BoardRecord) error {
	return fmt.Errorf("not implement")
}

// ReadBoardRecordFileRecord return boardRecord brd on index in record file.
func (db *DB) ReadBoardRecord(index uint) (*BoardRecord, error) {
	return nil, fmt.Errorf("not implement")
}

// RemoveBoardRecordFileRecord remove boardRecord brd on index in record file.
func (db *DB) RemoveBoardRecord(index uint) error {
	return fmt.Errorf("not implement")
}

// CreateArticleRecord returns new ArticleRecord with new filename in boardID and owner, date and title.
// This method will find a usable filename in board and occupy it.
func (db *DB) CreateArticleRecord(boardID, owner, date, title string) (ArticleRecord, error) {
	filename, err := db.connector.(WriteArticleConnector).CreateBoardArticleFilename(boardID)
	if err != nil {
		return nil, err
	}
	return db.connector.(WriteArticleConnector).NewArticleRecord(filename, owner, date, title)
}

// AddArticleRecordFileRecord append article ArticleRecord to boardID
func (db *DB) AddArticleRecordFileRecord(boardID string, article ArticleRecord) error {

	path, err := db.connector.GetBoardArticleRecordsPath(boardID)
	if err != nil {
		log.Println("bbs: open file error:", err)
		return err
	}
	log.Println("path:", path)

	return db.connector.(WriteArticleConnector).AddArticleRecordFileRecord(path, article)
}

// WriteBoardArticleFile writes content into filename in boardID
func (db *DB) WriteBoardArticleFile(boardID, filename string, content []byte) error {

	_, ok := db.connector.(WriteArticleConnector)
	if !ok {
		return fmt.Errorf("bbs: connector don't support WriteArticleConnector")
	}

	path, err := db.connector.GetBoardArticleFilePath(boardID, filename)
	if err != nil {
		log.Println("bbs: open file error:", err)
		return err
	}
	log.Println("path:", path)

	err = db.connector.(WriteArticleConnector).WriteBoardArticleFile(path, content)
	if err != nil {
		log.Println("bbs: write board article file error:", err)
		return err
	}
	return nil
}

func (db *DB) AppendBoardArticleFile(filename string, content []byte) error {
	c, ok := db.connector.(WriteArticleConnector)
	if !ok {
		return fmt.Errorf("bbs: connector don't support WriteArticleConnector")
	}
	err := c.AppendBoardArticleFile(filename, content)
	return err
}

// GetUserArticleRecordFile returns aritcle file which user posted.
func (db *DB) GetUserArticleRecordFile(userID string) ([]UserArticleRecord, error) {

	recs := []UserArticleRecord{}
	uac, ok := db.connector.(UserArticleConnector)
	if ok {

		path, err := uac.GetUserArticleRecordsPath(userID)
		if err != nil {
			log.Println("bbs: open file error:", err)
			return nil, err
		}
		log.Println("path:", path)

		recs, err = uac.ReadUserArticleRecordFile(path)
		if err != nil {
			log.Println("bbs: ReadUserArticleRecordFile error:", err)
			return nil, err
		}
		if len(recs) != 0 {
			return recs, nil
		}

	}

	boardRecords, err := db.ReadBoardRecords()
	if err != nil {
		log.Println("bbs: ReadBoardRecords error:", err)
		return nil, err
	}

	shouldSkip := func(boardID string) bool {
		if boardID == "ALLPOST" {
			return true
		}
		return false
	}

	for _, r := range boardRecords {
		if shouldSkip(r.BoardID()) {
			continue
		}

		ars, err := db.ReadBoardArticleRecordsFile(r.BoardID())
		if err != nil {
			log.Println("bbs: ReadBoardArticleRecordsFile error:", err)
			return nil, err
		}
		for _, ar := range ars {
			if ar.Owner() == userID {
				log.Println("board: ", r.BoardID(), len(recs))
				r := userArticleRecord{
					"board_id":   r.BoardID(),
					"title":      ar.Title(),
					"owner":      ar.Owner(),
					"article_id": ar.Filename(),
				}
				recs = append(recs, r)
			}
		}
	}

	return recs, nil
}

// GetUserCommentRecordFile returns the comment records of the specific user
// from all boards and all articles.
func (db *DB) GetUserCommentRecordFile(userID string) ([]UserCommentRecord, error) {

	recs := []UserCommentRecord{}
	ucc, ok := db.connector.(UserCommentConnector)
	if ok {
		path, err := ucc.GetUserCommentRecordsPath(userID)
		if err != nil {
			log.Println("bbs: open file error:", err)
			return nil, err
		}
		log.Println("path:", path)

		recs, err = ucc.ReadUserCommentRecordFile(path)
		if err != nil {
			log.Println("bbs: ReadUserCommentRecordFile error:", err)
			return nil, err
		}

		if len(recs) != 0 {
			return recs, nil
		}
	}

	// TODO: Implement a method to get the board records with the filter.
	// For example: db.ReadBoardRecordsFilter(skipBoardID []string)
	boardRecords, err := db.ReadBoardRecords()
	if err != nil {
		log.Println("bbs: ReadBoardRecords error:", err)
		return nil, err
	}

	shouldSkip := func(boardID string) bool {
		if boardID == "ALLPOST" {
			return true
		}
		return false
	}

	for _, r := range boardRecords {
		if shouldSkip(r.BoardID()) {
			continue
		}

		ucr, err := db.GetBoardUserCommentRecord(r.BoardID(), userID)
		if err != nil {
			log.Println("bbs: GetUserCommentRecordOfBoard error:", err)
			return nil, err
		}
		recs = append(recs, ucr...)
	}

	return recs, nil
}

// GetBoardUserCommentRecord returns the comment records of the user from the
// specific board.
func (db *DB) GetBoardUserCommentRecord(boardID, userID string) (recs []UserCommentRecord, err error) {

	ars, err := db.ReadBoardArticleRecordsFile(boardID)
	if err != nil {
		log.Println("bbs: ReadBoardArticleRecordsFile error:", err)
		return nil, err
	}

	for _, ar := range ars {
		crs, err := db.GetBoardArticleCommentRecords(boardID, ar)
		if err != nil {
			log.Println("bbs: GetBoardArticleCommentRecords error:", err)
			return nil, err
		}
		for _, cr := range crs {
			if userID != cr.Owner() {
				continue
			}
			recs = append(recs, cr)
		}
	}

	return recs, nil
}

// GetBoardArticleCommentRecords returns the comment records of the specific
// article.
func (db *DB) GetBoardArticleCommentRecords(boardID string, ar ArticleRecord) (crs []UserCommentRecord, err error) {

	content, err := db.ReadBoardArticleFile(boardID, ar.Filename())
	if err != nil {
		log.Println("bbs: ReadBoardArticleFile error:", err)
		return nil, err
	}

	floorCnt := uint32(1)
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		l := FilterStringANSI(scanner.Text())
		cr, err := NewUserCommentRecord(floorCnt, l, boardID, ar)
		if err != nil {
			// skip non-comment line
			if errors.Is(err, ErrNotUserComment) {
				continue
			}
			log.Println("bbs: NewUserCommentRecord error:", err)
			return nil, err
		}
		crs = append(crs, cr)
		floorCnt++
	}

	if len(crs) > 1 {
		log.Println(content)
	}

	return crs, nil
}

func (db *DB) GetUserDrafts(userID, draftID string) (UserDraft, error) {

	path, err := db.connector.(UserDraftConnector).GetUserDraftPath(userID, draftID)
	if err != nil {
		log.Println("bbs: GetUserDraftPath error:", err)
		return nil, err
	}
	log.Println("path:", path)

	raw, err := db.connector.(UserDraftConnector).ReadUserDraft(path)
	if err != nil {
		return nil, err
	}

	return NewUserDraft(raw), nil
}

func (db *DB) DeleteUserDraft(userID, draftID string) error {

	path, err := db.connector.(UserDraftConnector).GetUserDraftPath(userID, draftID)
	if err != nil {
		log.Println("bbs: GetUserDraftPath error:", err)
		return err
	}
	log.Println("path:", path)

	return db.connector.(UserDraftConnector).DeleteUserDraft(path)
}

func (db *DB) WriteUserDraft(userID, draftID string, draftContent userDraft) error {
	path, err := db.connector.(UserDraftConnector).GetUserDraftPath(userID, draftID)
	if err != nil {
		log.Println("bbs: GetUserDraftPath error:", err)
		return err
	}

	log.Println("path:", path)
	return db.connector.(UserDraftConnector).WriteUserDraft(path, draftContent.Raw())
}
