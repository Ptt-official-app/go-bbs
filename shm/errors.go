package shm

import "errors"

var (
	ErrInvalidShm        = errors.New("invalid shm")
	ErrUnableToAttachShm = errors.New("unable to attach shm")
	ErrUnableToCloseShm  = errors.New("unable to close shm")
)
