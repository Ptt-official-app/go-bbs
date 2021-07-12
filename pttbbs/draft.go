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

func (c *Connector) WriteUserDraft(filename string, draft []byte) error {
	return ioutil.WriteFile(filename, draft, 0o644)
}
