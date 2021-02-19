package cache

import (
	"fmt"
	"os"
	"runtime"
	// "strings"
)

type Mmap struct {
	file *os.File
	buf  []byte
}

func CreateMmap(filePath string, size int64) (*Mmap, error) {
	f, err := os.Create(filePath)
	// f, err := os.Open("../../../dump.shm")
	if err != nil {
		return nil, fmt.Errorf("create error: %w", err)
	}
	err = f.Truncate(size)
	if err != nil {
		return nil, fmt.Errorf("truncate error: %v", err)
	}

	return openFile(f)

}

func OpenMmap(filePath string) (*Mmap, error) {
	// := os.Open(filePath)
	f, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	// f, err := os.Open("../../../dump.shm")
	if err != nil {
		return nil, fmt.Errorf("open error: %w", err)
	}
	return openFile(f)
}

func openFile(f *os.File) (*Mmap, error) {

	fd := int(f.Fd())
	// fmt.Println("fd:", fd)

	stat, err := f.Stat()
	if err != nil {
		// fmt.Println("stat error", err)
		return nil, fmt.Errorf("stat error: %w", err)
	}

	size := int(stat.Size())
	// fmt.Println("size", size)

	b, err := openMmap(fd, size)
	if err != nil {
		fmt.Println("mmap error", err)
		return nil, fmt.Errorf("mmap error: %v", err)
	}

	ret := Mmap{
		file: f,
		buf:  b,
	}

	// For GC
	runtime.SetFinalizer(&ret, (*Mmap).Close)
	return &ret, nil
}

func (m *Mmap) Bytes() []byte {
	return m.buf
}

func (m *Mmap) Close() error {
	err := closeMmap(m.buf)

	m.file.Close()

	// release GC setting
	runtime.SetFinalizer(m, nil)
	return err

}

func RemoveMmap(filePath string) error {
	return os.Remove(filePath)
}
