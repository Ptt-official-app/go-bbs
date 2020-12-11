package bbs

import (
	"fmt"
)

func GetPasswdsPath(workDirectory string) (string, error) {
	return fmt.Sprintf("%s/.PASSWDS", workDirectory), nil
}

func GetBoardPath(workDirectory string) (string, error) {
	return fmt.Sprintf("%s/.BRD", workDirectory), nil
}
