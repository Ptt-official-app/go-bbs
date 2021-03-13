package bbs

// benchmark with go test -bench=.
/*
$ go test -bench=.
goos: linux
goarch: amd64
pkg: github.com/Ptt-official-app/go-bbs
cpu: Intel(R) Core(TM) i7-7500U CPU @ 2.70GHz
BenchmarkLevelDBAppend-4                      64          17167994 ns/op
BenchmarkRecordIOProtobufWrite-4             667           1821572 ns/op
BenchmarkRecordIOProtobufAppend-4            201           5656839 ns/op
BenchmarkProtobufWrite-4                     715           1970847 ns/op
BenchmarkProtobufAppend-4                    158           7869925 ns/op
BenchmarkProtobufBufWrite-4                  344           3447253 ns/op
BenchmarkProtobufArrayWrite-4               8366            148526 ns/op
BenchmarkProtobufArrayAppend-4                 5         221438740 ns/op
BenchmarkJSONStreamWrite-4                   562           2027456 ns/op
BenchmarkJSONStreamBufWrite-4               2785            451461 ns/op
BenchmarkJSONStreamAppend-4                  196           5495245 ns/op
BenchmarkJSONArrayWrite-4                    100          10640510 ns/op
BenchmarkJSONArrayAppend-4                     2        1415491800 ns/op
BenchmarkSqliteWrite-4                        54          19211187 ns/op
BenchmarkSqliteAppend-4                        1        8262016300 ns/op
PASS
ok      github.com/Ptt-official-app/go-bbs      34.153s
*/
import (
	"database/sql"
	"encoding/json"
	// "fmt"
	// "io"
	"io/ioutil"
	"os"
	"strconv"
	"testing"

	"github.com/eclesh/recordio"
	_ "github.com/mattn/go-sqlite3"
	"github.com/syndtr/goleveldb/leveldb"
	"google.golang.org/protobuf/proto"
)

var recordN = 1000

func BenchmarkLevelDBAppend(b *testing.B) {
	// LevelDB is the memory store backed by SSTable
	// It is useful in faster write and read
	// However, it is not good for close/reopen the db everytime
	a := ProtobufUserArticle{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}
	for i := 0; i < b.N; i++ {
		dir, err := ioutil.TempDir("", "leveldb")
		if err != nil {
			b.Errorf("%v", err)
			return
		}

		db, err := leveldb.OpenFile(dir, nil)
		if err != nil {
			b.Errorf("%v", err)
			return
		}

		for j := 0; j < recordN; j++ {
			out, err := proto.Marshal(&a)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			err = db.Put([]byte(strconv.Itoa(j)), out, nil)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
		}
		db.Close()
		os.RemoveAll(dir)
	}
}

func BenchmarkRecordIOProtobufWrite(b *testing.B) {
	// RecordIO is an append only data format
	// It is designed for faster sequential read/write
	a := ProtobufUserArticle{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}
	for i := 0; i < b.N; i++ {
		tmpfile, _ := ioutil.TempFile("./", "test_proto_buf")

		writer := recordio.NewWriter(tmpfile)

		for j := 0; j < recordN; j++ {
			out, err := proto.Marshal(&a)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			writer.Write(out)
		}
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}
}

func BenchmarkRecordIOProtobufAppend(b *testing.B) {
	a := ProtobufUserArticle{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}
	for i := 0; i < b.N; i++ {
		tmpfile, _ := ioutil.TempFile("./", "test_proto_buf")
		tmpfile.Close()

		for j := 0; j < recordN; j++ {
			file, err := os.OpenFile(tmpfile.Name(), os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			writer := recordio.NewWriter(file)
			out, err := proto.Marshal(&a)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			writer.Write(out)
			file.Close()
		}
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

func BenchmarkProtobufAppend(b *testing.B) {
	a := ProtobufUserArticle{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}
	for i := 0; i < b.N; i++ {
		// buf := []byte{}
		tmpfile, _ := ioutil.TempFile("./", "test_proto_buf")
		tmpfile.Close()

		for j := 0; j < recordN; j++ {
			file, err := os.OpenFile(tmpfile.Name(), os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				b.Errorf("%v", err)
				return
			}

			out, err := proto.Marshal(&a)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			file.Write(out)
			file.Close()
		}
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

		for j := 0; j < recordN; j++ {
			file, err := os.OpenFile(tmpfile.Name(), os.O_WRONLY, 0600)
			out, err := proto.Marshal(&a)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			buf = append(buf, out...)
			file.Close()
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

func BenchmarkProtobufArrayAppend(b *testing.B) {
	a := ProtobufUserArticle{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}
	for i := 0; i < b.N; i++ {
		// buf := []byte{}
		tmpfile, _ := ioutil.TempFile("./", "test_proto_buf")
		tmpfile.Close()

		for j := 0; j < recordN; j++ {
			file, err := os.OpenFile(tmpfile.Name(), os.O_RDWR, 0600)
			buf, err := ioutil.ReadAll(file)
			file.Seek(0, os.SEEK_SET)
			arr := ProtobufUserArticleList{}
			err = proto.Unmarshal(buf, &arr)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			arr.Items = append(arr.Items, &a)

			out, err := proto.Marshal(&arr)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			file.Write(out)
			file.Close()
		}
		// fi, _ := os.Stat(tmpfile.Name())
		// b.Logf("filesize: %v", fi.Size())
		os.Remove(tmpfile.Name())
	}
}

func BenchmarkJSONStreamWrite(b *testing.B) {

	s := struct {
		BoardID   string
		ArticleID string
	}{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}

	for i := 0; i < b.N; i++ {
		tmpfile, _ := ioutil.TempFile("./", "test_json")
		for j := 0; j < recordN; j++ {

			out, err := json.Marshal(&s)
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

func BenchmarkJSONStreamBufWrite(b *testing.B) {

	s := struct {
		BoardID   string
		ArticleID string
	}{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}

	for i := 0; i < b.N; i++ {
		tmpfile, _ := ioutil.TempFile("./", "test_json")
		buf := []byte{}
		for j := 0; j < recordN; j++ {

			out, err := json.Marshal(&s)
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

func BenchmarkJSONStreamAppend(b *testing.B) {

	s := struct {
		BoardID   string
		ArticleID string
	}{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}

	for i := 0; i < b.N; i++ {
		tmpfile, _ := ioutil.TempFile("./", "test_json")
		tmpfile.Close()
		for j := 0; j < recordN; j++ {
			file, err := os.OpenFile(tmpfile.Name(), os.O_APPEND|os.O_WRONLY, 0600)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			out, err := json.Marshal(&s)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			file.Write(out)
			file.Close()

		}

		// fi, _ := os.Stat(tmpfile.Name())
		// b.Logf("filesize: %v", fi.Size())
		os.Remove(tmpfile.Name())
	}
}

func BenchmarkJSONArrayWrite(b *testing.B) {
	type T struct {
		BoardID   string
		ArticleID string
	}
	arr := []T{}

	s := T{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}

	for i := 0; i < b.N; i++ {
		tmpfile, _ := ioutil.TempFile("./", "test_json")
		for j := 0; j < recordN; j++ {
			arr = append(arr, s)
		}

		out, err := json.Marshal(&arr)
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

func BenchmarkJSONArrayAppend(b *testing.B) {
	type T struct {
		BoardID   string
		ArticleID string
	}
	arr := []T{}

	s := T{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}

	for i := 0; i < b.N; i++ {
		tmpfile, _ := ioutil.TempFile("./", "test_json")
		tmpfile.Close()
		for j := 0; j < recordN; j++ {
			file, err := os.OpenFile(tmpfile.Name(), os.O_RDWR, 0600)
			buf, err := ioutil.ReadAll(file)
			file.Seek(0, os.SEEK_SET)

			err = json.Unmarshal(buf, &arr)
			arr = append(arr, s)
			out, err := json.Marshal(&arr)
			if err != nil {
				b.Errorf("%v", err)
				return
			}
			file.Write(out)
			file.Close()

		}
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
		tmpfile.Close()

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

		// fi, _ := os.Stat(tmpfile.Name())
		// b.Logf("filesize: %v", fi.Size())
		os.Remove(tmpfile.Name())

	}
}

func BenchmarkSqliteAppend(b *testing.B) {

	s := struct {
		BoardID   string
		ArticleID string
	}{
		BoardID:   "Soft_Job",
		ArticleID: "M.1610976994.A.2C8",
	}

	for i := 0; i < b.N; i++ {
		tmpfile, _ := ioutil.TempFile("./", "test_sqlite")
		tmpfile.Close()
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

		db.Close()

		for j := 0; j < recordN; j++ {

			db, err := sql.Open("sqlite3", tmpfile.Name())
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
			_, err = stmt.Exec(s.BoardID, s.ArticleID)
			if err != nil {
				b.Errorf("%v", err)
				return
			}

			tx.Commit()
			stmt.Close()
			db.Close()
		}

		// fi, _ := os.Stat(tmpfile.Name())
		// b.Logf("filesize: %v", fi.Size())
		os.Remove(tmpfile.Name())

	}
}
