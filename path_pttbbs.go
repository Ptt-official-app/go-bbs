// Copyright 2020 Pichu Chen, The PTT APP Authors

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This file manages the path rules of files in BBS system.

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

// Get Directory normal file path of board
func GetBoardArticleDirectoryath(workDirectory string, boardId string) (string, error) {
	return fmt.Sprintf("%s/boards/%c/%s/.DIR", workDirectory, boardId[0], boardId), nil
}

// Get Directory digest file path of board
func GetBoardNameFilePath(workDirectory string, boardId string) (string, error) {
	return fmt.Sprintf("%s/boards/%c/%s/.Name", workDirectory, boardId[0], boardId), nil
}
