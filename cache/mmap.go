package cache

import (
	"fmt"
	"os"
	"runtime"
	// "strings"
)

type Mmap struct {
	file *os.File
	Buf  []byte
}

func Open(filePath string) (*Mmap, error) {
	// filePath := strings.Replace(connectionString, "file:", "", 1)
	f, err := os.Open(filePath)
	// f, err := os.Open("../../../dump.shm")
	if err != nil {
		fmt.Println("open shm fail", err)
		return nil, fmt.Errorf("open error: %v", err)
	}
	fd := int(f.Fd())
	fmt.Println("fd:", fd)

	stat, err := f.Stat()
	if err != nil {
		fmt.Println("stat error", err)
		return nil, fmt.Errorf("stat error: %v", err)
	}

	size := int(stat.Size())
	fmt.Println("size", size)

	b, err := openMmap(fd, size)
	if err != nil {
		fmt.Println("mmap error", err)
		return nil, fmt.Errorf("mmap error: %v", err)
	}

	ret := Mmap{
		file: f,
		Buf:  b,
	}

	// For GC
	runtime.SetFinalizer(ret, (*Mmap).Close)
	return &ret, nil
}

func (m *Mmap) Close() {
	closeMmap(m.Buf)
	m.file.Close()

	// release GC setting
	runtime.SetFinalizer(m, nil)

}
