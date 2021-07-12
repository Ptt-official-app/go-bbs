package pttbbs

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Ptt-official-app/go-bbs"
)

type Connector struct {
	home string
}

func init() {
	bbs.Register("pttbbs", &Connector{})
}

// Open connect a file directory or SHMs, dataSourceName pointer to bbs home
// And it can append argument for SHM
// for example `file:///home/bbs/?UTMP=1993`
func (c *Connector) Open(dataSourceName string) error {
	if strings.HasPrefix(dataSourceName, "file://") {
		s := dataSourceName[len("file://"):]
		seg := strings.Split(s, "?")
		c.home = seg[0]
	} else {
		c.home = dataSourceName
	}
	return nil
}

func (c *Connector) GetUserRecordsPath() (string, error) {
	return GetPasswdsPath(c.home)
}

func (c *Connector) ReadUserRecordsFile(filename string) ([]bbs.UserRecord, error) {
	rec, err := OpenUserecFile(filename)
	ret := make([]bbs.UserRecord, len(rec))
	for i, v := range rec {
		ret[i] = v
	}
	return ret, err
}

func (c *Connector) GetUserDraftPath(userID, draftID string) (string, error) {
	return GetUserDraftPath(c.home, userID, draftID)
}

func (c *Connector) GetUserFavoriteRecordsPath(userID string) (string, error) {
	return GetUserFavoritePath(c.home, userID)
}

func (c *Connector) ReadUserFavoriteRecordsFile(filename string) ([]bbs.FavoriteRecord, error) {
	rec, err := OpenFavFile(filename)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: OpenFavFile error: %w", err)
	}

	bPath, err := c.GetBoardRecordsPath()
	if err != nil {
		return nil, fmt.Errorf("pttbbs: GetBoardRecordsPath error: %w", err)
	}
	br, err := OpenBoardHeaderFile(bPath)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: ReadBoardRecordsFile error: %w", err)
	}
	appendBoardID(rec.Folder, br)
	ret := make([]bbs.FavoriteRecord, len(rec.Folder.FavItems))
	for i, v := range rec.Folder.FavItems {
		ret[i] = v
	}
	return ret, err
}

func appendBoardID(folder *FavFolder, brd []*BoardHeader) {
	for _, item := range folder.FavItems {
		switch item.Item.(type) {
		case *FavBoardItem:
			bid := item.Item.(*FavBoardItem).BoardID - 1
			item.Item.(*FavBoardItem).boardID = brd[bid].BrdName
		case *FavFolderItem:
			appendBoardID(item.Item.(*FavFolderItem).ThisFolder, brd)
		case *FavLineItem:
			break
		default:
			// logger.Warningf("parseFavoriteFolderItem unknown favItem type")
		}
	}
}

func (c *Connector) GetBoardRecordsPath() (string, error) {
	return GetBoardPath(c.home)
}

func (c *Connector) ReadBoardRecordsFile(path string) ([]bbs.BoardRecord, error) {
	rec, err := OpenBoardHeaderFile(path)
	ret := make([]bbs.BoardRecord, len(rec))
	for i, v := range rec {
		ret[i] = v
	}
	return ret, err
}

func (c *Connector) GetBoardArticleRecordsPath(boardID string) (string, error) {
	return GetBoardArticlesDirectoryPath(c.home, boardID)
}

func (c *Connector) ReadArticleRecordsFile(filename string) ([]bbs.ArticleRecord, error) {
	var fileHeaders []*FileHeader
	var err error
	fileHeaders, err = OpenFileHeaderFile(filename)
	if err != nil {
		return nil, err
	}
	ret := make([]bbs.ArticleRecord, len(fileHeaders))
	for i, v := range fileHeaders {
		ret[i] = v
	}
	return ret, err
}

func (c *Connector) GetBoardTreasureRecordsPath(boardID string, treasureID []string) (string, error) {
	return GetBoardTreasuresDirectoryPath(c.home, boardID, treasureID)
}

func (c *Connector) GetBoardArticleFilePath(boardID string, filename string) (string, error) {
	return GetBoardArticleFilePath(c.home, boardID, filename)
}

func (c *Connector) GetBoardTreasureFilePath(boardID string, treasureID []string, filename string) (string, error) {
	return GetBoardTreasureFilePath(c.home, boardID, treasureID, filename)
}

// ReadBoardArticleFile returns raw file of specific filename article.
func (c *Connector) ReadBoardArticleFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: open file error: %v", err)
	}
	defer file.Close()
	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: readfile error: %v", err)
	}
	return buf, err
}
