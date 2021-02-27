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

// This file is for favorite function.

package pttbbs

import (
	"github.com/Ptt-official-app/go-bbs"

	"bytes"
	"encoding"
	"encoding/binary"
	"errors"
	"fmt"
	"io/ioutil"
	"time"
)

// For Current PTT
// Please see https://github.com/ptt/pttbbs/blob/master/include/fav.h
// https://github.com/ptt/pttbbs/blob/af507e0029c4e6b3a564ec98328ffe7cd7fd16be/mbbsd/fav.c
//
// This Fav parser parses favorites file. The favorites file contains the following
// 1. 2 bytes for FavFolder.Version
//
// Followed by at least 1 FavFolder, each FavFolder contains
// 2. 2 bytes for FavFolder.NBoards, how many boards in fav
// 3. 1 byte for FavFolder.NLines, how many line separator in fav
// 4. 1 byte for FavFolder.NFolders, how many folder in fav
// FavItemTypeBoard / FavItemTypeFolder / FavItemTypeLine is wrapped inside FavItem.Item and can be cast later.
// So the total items in this file will be (countOfItems = FavFolder.NBoards + FavFolder.NFolders + FavFolder.NLines)
// Followed by a list of FavItem, each FavItem pre-allocates:
//     FavItem itself takes 2 bytes for FavItemType and FavAttr
//     FavItemTypeBoard pre-allocates 12 bytes
//     FavItemTypeLine pre-allocates 1 bytes
//     FavItemTypeFolder pre-allocates 50 bytes
// Lastly followed by folders as another FavFolder

const (
	TIME4TBytes             = 4 // Bytes for time4_t
	favPreAlloc             = 8
	sizeOfPttFavBoardBytes  = 12 // Each FavBoardItem takes this many bytes
	sizeOfPttFavFolderBytes = 50 // Each FavFolderItem takes bytes
	sizeOfPttFavLineBytes   = 1  // Each FavLineItem takes bytes
)

var (
	InvalidFavTypeError  = errors.New("invalid Favorite type")
	IndexOutOfBoundError = errors.New("index out of range, file format invalid")
)

type FavItemType uint8

const (
	FavItemTypeBoard  FavItemType = 1 // FAVT_BOARD
	FavItemTypeFolder FavItemType = 2 // FAVT_FOLDER
	FavItemTypeLine   FavItemType = 3 // FAVT_LINE
)

// FavAttr represents fav attr
type FavAttr uint8

const (
	FavhFav    FavAttr = 0x00000001 // FAVH_FAV
	FavhTag    FavAttr = 0x00000002 // FAVH_TAG
	FavhUnread FavAttr = 0x00000004 // FAVH_UNREAD
	FavhAdmTag FavAttr = 0x00000008 // FAVH_ADM_TAG
)

// FavItem represents 1 Item in FavFolder
type FavItem struct {
	FavType FavItemType
	FavAttr uint8
	Item    interface{} // This could be either FavBoardItem / FavFolderItem / FavLineItem
}

func (f *FavItem) BoardID() string {
	if f.FavType != FavItemTypeBoard {
		return ""
	}
	return f.Item.(*FavBoardItem).boardID
}

func (f *FavItem) Title() string {
	if f.FavType == FavItemTypeLine {
		return "------------------------------------------"
	}
	if f.FavType == FavItemTypeFolder {
		return f.Item.(*FavFolderItem).Title
	}
	return ""
}

func (f *FavItem) Type() bbs.FavoriteType {
	switch f.FavType {
	case FavItemTypeBoard:
		return bbs.FavoriteTypeBoard
	case FavItemTypeFolder:
		return bbs.FavoriteTypeFolder
	case FavItemTypeLine:
		return bbs.FavoriteTypeLine
	}
	return bbs.FavoriteTypeBoard

}

func (f *FavItem) Records() []bbs.FavoriteRecord {
	if f.FavType != FavItemTypeFolder {
		return nil
	}
	rec := f.Item.(*FavFolderItem).ThisFolder.FavItems
	ret := make([]bbs.FavoriteRecord, len(rec))
	for i, v := range rec {
		ret[i] = v
	}
	return ret
}

// GetBoard tries to cast Item to FavBoardItem; return nil if it is not
func (favt *FavItem) GetBoard() *FavBoardItem {
	if ret, ok := favt.Item.(*FavBoardItem); ok {
		return ret
	}
	return nil
}

// GetFolder tries to cast Item to FavFolderItem; return nil if it is not
func (favt *FavItem) GetFolder() *FavFolderItem {
	if ret, ok := favt.Item.(*FavFolderItem); ok {
		return ret
	}
	return nil
}

// GetLine tries to cast Item to FavLineItem; return nil if it is not
func (favt *FavItem) GetLine() *FavLineItem {
	if ret, ok := favt.Item.(*FavLineItem); ok {
		return ret
	}
	return nil
}

// FavFile represents the entire fav file. Starts with 2 bytes of Version and at most 1 FavFolder.
type FavFile struct {
	Version uint16
	Folder  *FavFolder
}

// FavFolder represents a folder in .fav file. Each folder could contain NBoards of board, NLines of lines
// and NFolders of sub-folders.
type FavFolder struct {
	NAlloc   uint16
	DataTail uint16
	NBoards  uint16
	NLines   uint8
	NFolders uint8
	LineID   uint8
	FolderID uint8
	FavItems []*FavItem
}

// FavBoardItem represents a Board in FavFolder. FavBoardItem takes 12 bytes
type FavBoardItem struct {
	BoardID   uint32
	LastVisit time.Time
	Attr      uint32
	boardID   string
}

// FavFolderItem represents a Folder in FavFolder. FavFolderItem takes 50 bytes
type FavFolderItem struct {
	FolderID   uint8
	Title      string
	ThisFolder *FavFolder
}

// FavLineItem represents a Line in FavFolder. FavLineItem takes 1 byte
type FavLineItem struct {
	LineID uint8
}

// OpenFavFile reads a fav file
func OpenFavFile(filename string) (*FavFile, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return NewFavFile(data)
}

// NewFavFile parse data and return FavFile
func NewFavFile(data []byte) (*FavFile, error) {
	ret := &FavFile{}
	size := 2
	ret.Version = binary.LittleEndian.Uint16(data[0:size])

	var err error
	ret.Folder, _, err = NewFavFolder(data, size)
	if err != nil {
		return nil, err
	}
	return ret, err
}

// getDataNumber returns the count of total items in FavFolder
func (f *FavFolder) getDataNumber() uint16 {
	return f.NBoards + uint16(f.NFolders) + uint16(f.NLines)
}

// NewFavFolder takes a []byte, parse it starting with startIndex, return an instance of FavFolder, endIndex
// and error.
func NewFavFolder(data []byte, startIndex int) (*FavFolder, int, error) {
	// data must at least has 4 bytes for a new FavFolder
	if len(data) < startIndex+4 {
		return nil, startIndex, IndexOutOfBoundError
	}
	ret := &FavFolder{}
	c := startIndex // current index

	size := 2
	ret.NBoards = binary.LittleEndian.Uint16(data[c : c+size])
	c += size

	size = 1
	ret.NLines = data[c]
	c += size

	size = 1
	ret.NFolders = data[c]
	c += size

	ret.DataTail = ret.getDataNumber()
	ret.NAlloc = ret.DataTail + favPreAlloc
	ret.LineID = 0
	ret.FolderID = 0

	itemCount := ret.DataTail
	ret.FavItems = make([]*FavItem, itemCount)
	var err error

	// There are itemCount items, parse and insert them one by one
	for itemCount > 0 {
		n := len(ret.FavItems) - int(itemCount) // calculate index
		ret.FavItems[n], c, err = NewFavItem(data, c)
		if err != nil {
			return nil, 0, err
		}
		itemCount--
	}

	// Parse and insert next folder, if any
	for _, item := range ret.FavItems {
		if f, ok := item.Item.(*FavFolderItem); ok {
			var nextFolder *FavFolder
			nextFolder, c, err = NewFavFolder(data, c)
			if err != nil {
				return nil, c, err
			}
			ret.FolderID++
			f.FolderID = ret.FolderID
			f.ThisFolder = nextFolder
		}
		if f, ok := item.Item.(*FavLineItem); ok {
			ret.LineID++
			f.LineID = ret.LineID
		}
	}

	return ret, c, nil
}

// NewFavItem parse data starting from startIndex and return FavItem. FavItem.Item might be either FavBoardItem,
// FavFolderItem or FavLineItem
func NewFavItem(data []byte, startIndex int) (*FavItem, int, error) {
	// data at least must have 2 bytes for a new FavItem
	if len(data) < startIndex+2 {
		return nil, startIndex, IndexOutOfBoundError
	}
	ret := &FavItem{}
	c := startIndex // current index

	size := 1
	ret.FavType = FavItemType(data[c])
	c += size

	size = 1
	ret.FavAttr = data[c]
	c += size

	var err error
	var item interface{}

	switch ret.FavType {
	case FavItemTypeBoard:
		item, c, err = NewFavBoardItem(data, c)
	case FavItemTypeLine:
		item, c, err = NewFavLineItem(data, c)
	case FavItemTypeFolder:
		item, c, err = NewFavFolderItem(data, c)
	default:
		err = InvalidFavTypeError
	}
	if err != nil {
		return nil, c, err
	}
	ret.Item = item

	return ret, c, err
}

// NewFavBoardItem takes a []byte and parse it starting from startIndex, return FavBoardItem, end index and error
func NewFavBoardItem(data []byte, startIndex int) (*FavBoardItem, int, error) {
	if len(data) < startIndex+sizeOfPttFavBoardBytes {
		return nil, startIndex, IndexOutOfBoundError
	}
	ret := &FavBoardItem{}
	c := startIndex

	size := 4
	ret.BoardID = binary.LittleEndian.Uint32(data[c : c+size])
	c += size

	size = TIME4TBytes // use 4 bytes for time.Time
	ret.LastVisit = time.Unix(int64(binary.LittleEndian.Uint32(data[c:c+size])), 0)
	c += size

	// This attr is a char in fav.h which should have been 1 byte. However, from the sample file
	// we can see a Board takes 12 bytes, 4 bytes for BoardID, 4 bytes for LastVisit, so allocate the remaining
	// 4 byte to attr. May need double check on this.
	size = 4
	ret.Attr = binary.LittleEndian.Uint32(data[c : c+size])
	c += size

	return ret, c, nil
}

// NewFavFolderItem takes a []byte and parse it starting from startIndex, return FavFolderItem, end index and error
func NewFavFolderItem(data []byte, startIndex int) (*FavFolderItem, int, error) {
	if len(data) < startIndex+sizeOfPttFavFolderBytes {
		return nil, startIndex, IndexOutOfBoundError
	}
	ret := &FavFolderItem{}
	c := startIndex

	size := 1
	ret.FolderID = data[c]
	c += size

	size = PTT_BTLEN + 1
	ret.Title = big5uaoToUTF8String(bytes.Split(data[c:c+size], []byte("\x00"))[0])
	c += size

	return ret, c, nil
}

// NewFavLineItem takes a []byte and parse it starting from startIndex, return FavLineItem, end index and error
func NewFavLineItem(data []byte, startIndex int) (*FavLineItem, int, error) {
	if len(data) < startIndex+sizeOfPttFavLineBytes {
		return nil, startIndex, IndexOutOfBoundError
	}
	ret := &FavLineItem{}
	c := startIndex

	ret.LineID = data[c]
	c++
	return ret, c, nil
}

func (favf *FavFile) MarshalBinary() ([]byte, error) {
	ret := make([]byte, 2)

	binary.LittleEndian.PutUint16(ret[0:2], favf.Version)
	folderInBytes, err := favf.Folder.MarshalBinary()
	if err != nil {
		return nil, err
	}
	ret = append(ret, folderInBytes...)

	return ret, nil
}

func (favf *FavFolder) MarshalBinary() ([]byte, error) {
	ret := make([]byte, 4)
	c := 0

	size := 2
	binary.LittleEndian.PutUint16(ret[c:c+size], favf.NBoards)
	c += size

	ret[c] = favf.NLines
	c += 1

	ret[c] = favf.NFolders

	for _, item := range favf.FavItems {
		encoded, err := item.MarshalBinary()
		if err != nil {
			return nil, err
		}
		ret = append(ret, encoded...)
	}

	for _, item := range favf.FavItems {
		if f, ok := item.Item.(*FavFolderItem); ok {
			encoded, err := f.ThisFolder.MarshalBinary()
			if err != nil {
				return nil, err
			}
			ret = append(ret, encoded...)
		}
	}

	return ret, nil
}

func (favi *FavItem) MarshalBinary() ([]byte, error) {
	ret := make([]byte, 2)

	ret[0] = uint8(favi.FavType)
	ret[1] = favi.FavAttr
	favim, ok := favi.Item.(encoding.BinaryMarshaler)
	if !ok {
		return nil, fmt.Errorf("FavItem.Item must implement encoding.BinaryMarshaler")
	}
	encoded, err := favim.MarshalBinary()
	if err != nil {
		return nil, err
	}
	ret = append(ret, encoded...)

	return ret, nil
}

func (favbi *FavBoardItem) MarshalBinary() ([]byte, error) {
	ret := make([]byte, sizeOfPttFavBoardBytes)
	c := 0

	size := 4
	binary.LittleEndian.PutUint32(ret[c:c+size], favbi.BoardID)
	c += size

	binary.LittleEndian.PutUint32(ret[c:c+size], uint32(favbi.LastVisit.Unix()))
	c += size

	binary.LittleEndian.PutUint32(ret[c:c+size], favbi.Attr)

	return ret, nil
}

func (favfi *FavFolderItem) MarshalBinary() ([]byte, error) {
	ret := make([]byte, sizeOfPttFavFolderBytes)
	ret[0] = favfi.FolderID

	size := PTT_BTLEN + 1
	copy(ret[1:1+size], utf8ToBig5UAOString(favfi.Title))

	return ret, nil
}

func (favli *FavLineItem) MarshalBinary() ([]byte, error) {
	ret := make([]byte, sizeOfPttFavLineBytes)
	ret[0] = favli.LineID
	return ret, nil
}
