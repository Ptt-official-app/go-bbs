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
	// specified in pttstruct.h: line 51 (len(content) + 1)
	PASSLEN = 14
)

var (
	ErrInvalidCrypt = errors.New("invalid crypt")
	mu              sync.Mutex
)

//Fcrypt
//Params
//	key: the input-key (input-passwd) to be encrypted / checked
//	salt: the salt (expected-passwd-hash) in crypt(3)
//
//Return
//	[]byte: encrypted passwd, should be the same as salt if salt is the expected-passwd-hash.
//  error: err
func Fcrypt(key []byte, salt []byte) ([]byte, error) {
	// cpasswd is static unsigned char buff[20], requiring mutex.
	mu.Lock()
	defer mu.Unlock()

	ckey := C.CBytes(key)
	defer C.free(ckey)

	csalt := C.CBytes(salt)
	defer C.free(csalt)

	// cpasswd is static unsigned char buff[20] in bbscrypt.c: line 543, no need to free.
	cpasswdHash, err := C.fcrypt_wrapper(ckey, csalt)
	if err != nil {
		return nil, err
	}

	passwdHash := C.GoBytes(cpasswdHash, PASSLEN)
	// specified in bbscrypt.c: line 592
	passwdHash[PASSLEN-1] = 0
	return passwdHash, nil
}
