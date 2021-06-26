package bbs

import (
	"encoding/hex"
	"strings"
	"testing"
	"time"
)

func TestGetBoardArticleCommentRecords(t *testing.T) {

	anyBoardID := ""
	anyArticleRecord := &MockArticleRecord{}
	articleContent := `
a740 aacc 3a20 5359 534f 5020 28af ab29
20ac ddaa 4f3a 2074 6573 740a bcd0 c344
3a20 5bb0 ddc3 445d 2054 6869 7320 506f
7374 2077 696c 6c20 6265 2069 6e20 6d61
6e0a aec9 b6a1 3a20 5361 7420 4a61 6e20
2039 2031 393a 3032 3a31 3220 3230 3231
0a0a 0a54 6573 740a 0a2d 2d0a a1b0 20b5
6fab 48af b83a 20b7 73a7 e5bd f0bd f028
7074 7432 2e63 6329 2c20 a8d3 a6db 3a20
312e 3136 342e 3131 312e 3135 360a a1b0
20a4 e5b3 b9ba f4a7 7d3a 2068 7474 703a
2f2f 7777 772e 7074 742e 6363 2f62 6273
2f74 6573 742f 4d2e 3136 3130 3231 3839
3334 2e41 2e46 4130 2e68 746d 6c0a 1b5b
313b 3331 6da1 f720 1b5b 3333 6d53 5953
4f50 1b5b 6d1b 5b33 336d 3a50 7573 6820
2020 2020 2020 2020 2020 2020 2020 2020
2020 2020 2020 2020 2020 2020 2020 2020
2020 2020 2020 2020 2020 2020 2020 2020
2020 1b5b 6db1 c020 3031 2f30 3920 3139
3a30 320a 1b5b 313b 3331 6da1 f720 1b5b
3333 6d53 5953 4f50 1b5b 6d1b 5b33 336d
3a74 6573 7420 2020 2020 2020 2020 2020
2020 2020 2020 2020 2020 2020 2020 2020
2020 2020 2020 2020 2020 2020 2020 2020
2020 2020 2020 2020 1b5b 6db1 c020 3031
2f31 3720 3131 3a33 360a a1b0 201b 5b31
3b33 326d 7069 6368 7532 1b5b 303b 3332
6d3a c2e0 bffd a6dc acdd aa4f 2053 5953
4f50 1b5b 6d20 2020 2020 2020 2020 2020
2020 2020 2020 2020 2020 2020 2020 2020
2020 2020 2031 3131 2e32 3438 2e34 372e
3136 3420 3031 2f32 370a`

	returnPerfectBoardArticleFilePath := func() (string, error) {
		return anyBoardID, nil
	}
	returnPerfectReadBoardArticleFile := func() ([]byte, error) {
		return hexToByte(articleContent), nil
	}

	type fields struct {
		connector Connector
	}
	type args struct {
		boardID       string
		articleRecord ArticleRecord
	}
	type expected struct {
		recordCount int
		hasErr      bool
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		expected expected
	}{
		{
			name: "parse an article contains two user comments",
			fields: fields{
				connector: &fakeConnector{
					fakeGetBoardArticleFilePath: returnPerfectBoardArticleFilePath,
					fakeReadBoardArticleFile:    returnPerfectReadBoardArticleFile,
				},
			},
			args:     args{anyBoardID, anyArticleRecord},
			expected: expected{2, false},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{connector: tt.fields.connector}

			got, err := db.GetBoardArticleCommentRecords(tt.args.boardID, tt.args.articleRecord)

			if len(got) != tt.expected.recordCount {
				t.Errorf("record count = %v, expected %v\n", len(got), tt.expected.recordCount)
			}

			if (err != nil) != tt.expected.hasErr {
				t.Errorf("err = %v, should be nil\n", err)
			}
		})
	}

}

type fakeConnector struct {
	fakeOpen                        func() error
	fakeGetUserRecordsPath          func() (string, error)
	fakeReadUserRecordsFile         func() ([]UserRecord, error)
	fakeGetUserFavoriteRecordsPath  func() (string, error)
	fakeReadUserFavoriteRecordsFile func() ([]FavoriteRecord, error)
	fakeGetBoardRecordsPath         func() (string, error)
	fakeReadBoardRecordsFile        func() ([]BoardRecord, error)
	fakeGetBoardArticleRecordsPath  func() (string, error)
	fakeGetBoardTreasureRecordsPath func() (string, error)
	fakeReadArticleRecordsFile      func() ([]ArticleRecord, error)
	fakeGetBoardArticleFilePath     func() (string, error)
	fakeGetBoardTreasureFilePath    func() (string, error)
	fakeReadBoardArticleFile        func() ([]byte, error)
}

var _ Connector = &fakeConnector{}

func (c *fakeConnector) Open(dataSourceName string) error {
	return c.fakeOpen()
}

func (c *fakeConnector) GetUserRecordsPath() (string, error) {
	return c.fakeGetUserRecordsPath()
}

func (c *fakeConnector) ReadUserRecordsFile(name string) ([]UserRecord, error) {
	return c.fakeReadUserRecordsFile()
}

func (c *fakeConnector) GetUserFavoriteRecordsPath(userID string) (string, error) {
	return c.fakeGetUserFavoriteRecordsPath()
}

func (c *fakeConnector) ReadUserFavoriteRecordsFile(name string) ([]FavoriteRecord, error) {
	return c.fakeReadUserFavoriteRecordsFile()
}

func (c *fakeConnector) GetBoardRecordsPath() (string, error) {
	return c.fakeGetBoardRecordsPath()
}

func (c *fakeConnector) ReadBoardRecordsFile(name string) ([]BoardRecord, error) {
	return c.fakeReadBoardRecordsFile()
}

func (c *fakeConnector) GetBoardArticleRecordsPath(boardID string) (string, error) {
	return c.fakeGetBoardArticleRecordsPath()
}

func (c *fakeConnector) GetBoardTreasureRecordsPath(boardID string, treasureID []string) (string, error) {
	return c.fakeGetBoardTreasureRecordsPath()
}

func (c *fakeConnector) ReadArticleRecordsFile(name string) ([]ArticleRecord, error) {
	return c.fakeReadArticleRecordsFile()
}

func (c *fakeConnector) GetBoardArticleFilePath(boardID string, filename string) (string, error) {
	return c.fakeGetBoardArticleFilePath()
}

func (c *fakeConnector) GetBoardTreasureFilePath(boardID string, treasureID []string, name string) (string, error) {
	return c.fakeGetBoardTreasureFilePath()
}

func (c *fakeConnector) ReadBoardArticleFile(name string) ([]byte, error) {
	return c.fakeReadBoardArticleFile()
}

func hexToByte(input string) []byte {
	s := strings.ReplaceAll(input, " ", "")
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	b, _ := hex.DecodeString(s)
	return b
}

type MockArticleRecord struct {
	filename  string
	modified  time.Time
	recommend int8
	owner     string
	date      string
	title     string
	money     int
}

func (f *MockArticleRecord) Filename() string                  { return f.filename }
func (f *MockArticleRecord) Modified() time.Time               { return f.modified }
func (f *MockArticleRecord) SetModified(newModified time.Time) { f.modified = newModified }
func (f *MockArticleRecord) Recommend() int                    { return int(f.recommend) }
func (f *MockArticleRecord) Owner() string                     { return f.owner }
func (f *MockArticleRecord) Date() string                      { return f.date }
func (f *MockArticleRecord) Title() string                     { return f.title }
func (f *MockArticleRecord) Money() int                        { return f.money }
