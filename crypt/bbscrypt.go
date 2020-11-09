package crypt

// #cgo CFLAGS: -Os -Wno-missing-field-initializers -pipe -I./include -Wno-parentheses-equality
// #include "bbscrypt.h"
//
// void *fcrypt_wrapper(void *buf, void *salt) {
//   return fcrypt((char *)buf, (char *)salt);
// }
import "C"
import (
	"errors"
	"sync"
)

const (
	// specified in bbscrypt.c: line 594
	PASSWD_LEN = 13
)

var (
	ErrInvalidCrypt = errors.New("invalid crypt")
	mu              sync.Mutex
)

func Fcrypt(input []byte, expected []byte) ([]byte, error) {
	// cpasswd is static unsigned char buff[20], requiring mutex.
	mu.Lock()
	defer mu.Unlock()

	cinput := C.CBytes(input)
	defer C.free(cinput)

	cexpected := C.CBytes(expected)
	defer C.free(cexpected)

	// cpasswd is static unsigned char buff[20] in bbscrypt.c: line 543, no need to free.
	cpasswd, err := C.fcrypt_wrapper(cinput, cexpected)
	if err != nil {
		return nil, err
	}

	passwd := C.GoBytes(cpasswd, PASSWD_LEN)
	return passwd, nil
}
