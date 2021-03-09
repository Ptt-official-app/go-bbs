package bbs

// benchmark with go test -bench=.
/*
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/Ptt-official-app/go-bbs
BenchmarkProtobufWrite-4        	     100	  12488268 ns/op
BenchmarkProtobufArrayWrite-4   	     667	   1519035 ns/op
BenchmarkSqliteWrite-4          	      79	  16020372 ns/op
PASS
ok  	github.com/Ptt-official-app/go-bbs	4.517s
*/
import (
	"database/sql"
	"io/ioutil"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/protobuf/proto"
)

var recordN = 1000

func BenchmarkProtobufWrite(b *testing.B) {
	a := ProtobufUserArticle{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}
	for i := 0; i < b.N; i++ {
		// buf := []byte{}
		tmpfile, _ := ioutil.TempFile("./", "test_proto_buf")

		recordN := 1000
		for j := 0; j < recordN; j++ {

			out, err := proto.Marshal(&a)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			tmpfile.Write(out)
		}
		tmpfile.Close()
		// fi, _ := os.Stat(tmpfile.Name())
		// b.Logf("filesize: %v", fi.Size())
		os.Remove(tmpfile.Name())
	}
}

func BenchmarkProtobufArrayWrite(b *testing.B) {
	a := ProtobufUserArticle{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}
	for i := 0; i < b.N; i++ {
		// buf := []byte{}
		tmpfile, _ := ioutil.TempFile("./", "test_proto_buf")

		recordN := 1000

		items := []*ProtobufUserArticle{}
		for j := 0; j < recordN; j++ {
			items = append(items, &a)
		}
		arr := ProtobufUserArticleList{
			Items: items,
		}

		out, err := proto.Marshal(&arr)
		if err != nil {
			b.Errorf("%v", err)
			return
		}
		tmpfile.Write(out)
		tmpfile.Close()
		// fi, _ := os.Stat(tmpfile.Name())
		// b.Logf("filesize: %v", fi.Size())
		os.Remove(tmpfile.Name())
	}
}

func BenchmarkSqliteWrite(b *testing.B) {

	s := struct {
		BoardID   string
		ArticleID string
	}{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}

	for i := 0; i < b.N; i++ {
		tmpfile, _ := ioutil.TempFile("./", "test_sqlite")

		db, err := sql.Open("sqlite3", tmpfile.Name())
		if err != nil {
			b.Errorf("%v", err)
			return
		}
		creatStmt := `CREATE TABLE records (board_id string, article_id string);`
		_, err = db.Exec(creatStmt)
		if err != nil {
			b.Errorf("%v", err)
			return
		}
		tx, err := db.Begin()
		if err != nil {
			b.Errorf("%v", err)
			return
		}
		stmt, err := tx.Prepare("INSERT INTO records(board_id, article_id) VALUES (?,?)")
		if err != nil {
			b.Errorf("%v", err)
			return
		}
		for j := 0; j < recordN; j++ {
			_, err = stmt.Exec(s.BoardID, s.ArticleID)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
		}
		tx.Commit()
		stmt.Close()
		db.Close()

		tmpfile.Close()
		// fi, _ := os.Stat(tmpfile.Name())
		// b.Logf("filesize: %v", fi.Size())
		os.Remove(tmpfile.Name())

	}
}
