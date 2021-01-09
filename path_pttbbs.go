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
// We introduce global files at first, and path of users, path of
// normal article, path of treasure (man in maple system), path of
// digest.

package bbs

import (
	"fmt"
	"strings"
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

// Get Directory digest file path of board, `workDirectory` is BBSHome usually, `userId` means
// which user, and filename is actual file in user, notice that, we will check is file exist or
// user has permission for open this file. eg, .DIR file, so please check filename before call
// this function.
func GetUserMailPath(workDirectory string, userId string, filename string) (string, error) {
	return fmt.Sprintf("%s/home/%c/%s/%s", workDirectory, userId[0], userId, filename), nil
}

// Get Login Recent file path of user
func GetLoginRecentPath(workDirectory string, userid string) (string, error) {
	return fmt.Sprintf("%s/home/%c/%s/logins.recent", workDirectory, userid[0], userid), nil
}

// Get Directory normal file path of board
func GetBoardArticleDirectoryPath(workDirectory string, boardId string) (string, error) {
	return fmt.Sprintf("%s/boards/%c/%s/.DIR", workDirectory, boardId[0], boardId), nil
}

// Get Directory digest file path of board, `workDirectory` is BBSHome usually, `boardId` means
// which board, and filename is actual file in board, notice that, we will check is file exist or
// user has permission for open this file. eg, .DIR file, so please check filename before call
// this function.
func GetBoardArticlePath(workDirectory string, boardId string, filename string) (string, error) {
	return fmt.Sprintf("%s/boards/%c/%s/%s", workDirectory, boardId[0], boardId, filename), nil
}

// GetBoardTreasuresDirectoryPath return dir file path of specific board and path,
// `workDirectory` is BBSHome usually, `boardId` means which board and `path` is a slice
// figure out each directory, eg: `["M.971228479.A", "M.1035338027.A"]` in formosa BBS or
// `["D690", "D6C2", "D6D1"]` in pttbbs.
func GetBoardTreasuresDirectoryPath(workDirectory string, boardId string, path []string) (string, error) {
	path = append(path, "")
	subPath := strings.Join(path, "/")
	return fmt.Sprintf("%s/man/boards/%c/%s/%s.DIR", workDirectory, boardId[0], boardId, subPath), nil
}

func GetBoardTreasuresFilePath(workDirectory string, boardId string, path []string, filename string) (string, error) {
	path = append(path, "")
	subPath := strings.Join(path, "/")
	return fmt.Sprintf("%s/man/boards/%c/%s/%s%s", workDirectory, boardId[0], boardId, subPath, filename), nil
}

// Get Directory digest file path of board
func GetBoardNameFilePath(workDirectory string, boardId string) (string, error) {
	return fmt.Sprintf("%s/boards/%c/%s/.Name", workDirectory, boardId[0], boardId), nil
}
