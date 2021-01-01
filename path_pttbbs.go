package bbs

import (
	"fmt"
)

// Get Passwd file path of system
func GetPasswdsPath(workDirectory string) (string, error) {
	return fmt.Sprintf("%s/.PASSWDS", workDirectory), nil
}

// Get Board file path of system
func GetBoardPath(workDirectory string) (string, error) {
	return fmt.Sprintf("%s/.BRD", workDirectory), nil
}

// Get Favorite file path of user
func GetUserFavoritePath(workDirectory string, userid string) (string, error) {
	return fmt.Sprintf("%s/home/%c/%s/.fav", workDirectory, userid[0], userid), nil
}

// Get Login Recent file path of user
func GetLoginRecentPath(workDirectory string, userid string) (string, error) {
	return fmt.Sprintf("%s/home/%c/%s/logins.recent", workDirectory, userid[0], userid), nil
}
