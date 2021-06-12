package pttbbs

import (
	"io/ioutil"
	"os"
)

func (c *Connector) ReadUserDraft(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (c *Connector) DeleteUserDraft(filename string) error {
	return os.Remove(filename)
}
