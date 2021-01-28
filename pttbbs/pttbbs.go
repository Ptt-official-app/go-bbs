package pttbbs

import (
	"github.com/PichuChen/go-bbs"

	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func (c *Connector) GetUserFavoriteRecordsPath(userId string) (string, error) {
	return GetUserFavoritePath(c.home, userId)
}

func (c *Connector) ReadUserFavoriteRecordsFile(filename string) ([]bbs.FavoriteRecord, error) {
	rec, err := OpenFavFile(filename)

	bPath, err := c.GetBoardRecordsPath()
	if err != nil {
		return nil, fmt.Errorf("pttbbs: GetBoardRecordsPath error: %v", err)
	}
	br, err := OpenBoardHeaderFile(bPath)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: ReadBoardRecordsFile error: %v", err)
	}
	appendBoardId(rec.Folder, br)
	ret := make([]bbs.FavoriteRecord, len(rec.Folder.FavItems))
	for i, v := range rec.Folder.FavItems {
		ret[i] = v
	}
	return ret, err
}

func appendBoardId(folder *FavFolder, brd []*BoardHeader) {
	for _, item := range folder.FavItems {

		switch item.Item.(type) {
		case *FavBoardItem:
			bid := item.Item.(*FavBoardItem).BoardId - 1
			item.Item.(*FavBoardItem).boardId = brd[bid].BrdName
		case *FavFolderItem:
			appendBoardId(item.Item.(*FavFolderItem).ThisFolder, brd)
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

func (c *Connector) GetBoardArticleRecordsPath(boardId string) (string, error) {
	return GetBoardArticlesDirectoryPath(c.home, boardId)
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

func (c *Connector) GetBoardTreasureRecordsPath(boardId string, treasureId []string) (string, error) {
	return GetBoardTreasuresDirectoryPath(c.home, boardId, treasureId)
}

func (c *Connector) GetBoardArticleFilePath(boardId string, filename string) (string, error) {
	return GetBoardArticleFilePath(c.home, boardId, filename)
}

func (c *Connector) GetBoardTreasureFilePath(boardId string, treasureId []string, filename string) (string, error) {
	return GetBoardTreasureFilePath(c.home, boardId, treasureId, filename)
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