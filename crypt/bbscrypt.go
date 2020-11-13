package crypt

import (
	"errors"
)

var (
	ErrInvalidCrypt = errors.New("invalid crypt")
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
	passwdHash := [PASSLEN]byte{}
	cFcrypt(key, salt, &passwdHash)
	return passwdHash[:], nil
}
