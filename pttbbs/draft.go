package pttbbs

import (
	"os"
)

func (c *Connector) DeleteUserDraft(filename string) error {
	return os.Remove(filename)
}
