package bbs

// benchmark with go test -bench=.
/*
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: github.com/Ptt-official-app/go-bbs
BenchmarkSSTableProtobufWrite-4   	      48	  33028588 ns/op
BenchmarkProtobufWrite-4          	     100	  16207969 ns/op
BenchmarkProtobufBufWrite-4       	     710	   2152707 ns/op
BenchmarkProtobufArrayWrite-4     	     825	   1801981 ns/op
BenchmarkSqliteWrite-4            	      72	  17835609 ns/op
PASS
ok  	github.com/Ptt-official-app/go-bbs	10.021s
*/
import (
	"database/sql"
	"io/ioutil"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/thomasjungblut/go-sstables/recordio"
	"google.golang.org/protobuf/proto"
)

var recordN = 1000

func BenchmarkSSTableProtobufWrite(b *testing.B) {
	a := ProtobufUserArticle{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}
	for i := 0; i < b.N; i++ {
		// buf := []byte{}
		tmpfile, _ := ioutil.TempFile("./", "test_proto_buf")
		tmpfile.Close()

		writer, err := recordio.NewCompressedProtoWriterWithPath(tmpfile.Name(), recordio.CompressionTypeSnappy)
		if err != nil {
			b.Errorf("error: %v", err)
		}

		err = writer.Open()
		if err != nil {
			b.Errorf("error: %v", err)
		}

		for j := 0; j < recordN; j++ {

			_, err := writer.Write(&a)
			if err != nil {
				b.Errorf("error: %v", err)
			}
			// b.Logf("wrote a record at offset of %d bytes", recordOffset)
		}

		err = writer.Close()
		if err != nil {
			b.Errorf("error: %v", err)
		}
		// fi, _ := os.Stat(tmpfile.Name())
		// b.Logf("filesize: %v", fi.Size())
		os.Remove(tmpfile.Name())
	}
}

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

func BenchmarkProtobufBufWrite(b *testing.B) {
	a := ProtobufUserArticle{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}
	for i := 0; i < b.N; i++ {
		buf := []byte{}
		tmpfile, _ := ioutil.TempFile("./", "test_proto_buf")

		recordN := 1000
		for j := 0; j < recordN; j++ {

			out, err := proto.Marshal(&a)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			buf = append(buf, out...)
		}
		tmpfile.Write(buf)
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
