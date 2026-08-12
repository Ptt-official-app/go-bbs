package pttbbs

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/Ptt-official-app/go-bbs"
)

const (
	pttFavoriteVersion   = 3363
	pttFavoriteMaxItems  = 1024
	pttFavoriteMaxLines  = 64
	pttFavoriteMaxFolder = 64
)

var favoriteWriteMu sync.Mutex

// AddUserFavorite appends an item to the root .fav folder and persists the
// whole favorite tree using PTT's write-temp-then-rename strategy.
func (c *Connector) AddUserFavorite(userID string, options bbs.FavoriteCreateOptions) (bbs.FavoriteRecord, error) {
	if userID == "" || filepath.Base(userID) != userID || strings.ContainsAny(userID, `/\`) {
		return nil, fmt.Errorf("pttbbs: invalid user id %q", userID)
	}

	favoriteWriteMu.Lock()
	defer favoriteWriteMu.Unlock()

	path, err := c.GetUserFavoriteRecordsPath(userID)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: get favorite path: %w", err)
	}
	fav, err := loadFavoriteForWrite(path)
	if err != nil {
		return nil, err
	}
	if countFavoriteItems(fav.Folder) >= pttFavoriteMaxItems {
		return nil, fmt.Errorf("pttbbs: favorite item limit reached")
	}

	var item *FavItem
	switch options.Type {
	case bbs.FavoriteTypeLine:
		item, err = appendFavoriteLine(fav.Folder)
	case bbs.FavoriteTypeFolder:
		item, err = appendFavoriteFolder(fav.Folder, options.Title)
	case bbs.FavoriteTypeBoard:
		item, err = c.appendFavoriteBoard(fav.Folder, options.BoardID)
	default:
		err = fmt.Errorf("pttbbs: unsupported favorite type %d", options.Type)
	}
	if err != nil {
		return nil, err
	}

	fav.Version = pttFavoriteVersion
	data, err := fav.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("pttbbs: marshal favorite file: %w", err)
	}
	if err := replaceFavoriteFile(path, data); err != nil {
		return nil, err
	}
	return item, nil
}

func loadFavoriteForWrite(path string) (*FavFile, error) {
	data, err := os.ReadFile(path) // #nosec G304 -- path comes from the BBS home/user path helper.
	if err != nil {
		if os.IsNotExist(err) {
			return newEmptyFavoriteFile(), nil
		}
		return nil, fmt.Errorf("pttbbs: read favorite file: %w", err)
	}
	if len(data) == 0 {
		return newEmptyFavoriteFile(), nil
	}
	fav, err := UnmarshalFavFile(data)
	if err != nil {
		return nil, fmt.Errorf("pttbbs: parse favorite file: %w", err)
	}
	return fav, nil
}

func newEmptyFavoriteFile() *FavFile {
	return &FavFile{
		Version: pttFavoriteVersion,
		Folder: &FavFolder{
			NAlloc:   favPreAlloc,
			FavItems: []*FavItem{},
		},
	}
}

func countFavoriteItems(folder *FavFolder) int {
	if folder == nil {
		return 0
	}
	count := len(folder.FavItems)
	for _, item := range folder.FavItems {
		if child, ok := item.Item.(*FavFolderItem); ok {
			count += countFavoriteItems(child.ThisFolder)
		}
	}
	return count
}

func appendFavoriteLine(folder *FavFolder) (*FavItem, error) {
	if folder == nil {
		return nil, fmt.Errorf("pttbbs: favorite root folder is nil")
	}
	if folder.NLines >= pttFavoriteMaxLines {
		return nil, fmt.Errorf("pttbbs: favorite line limit reached")
	}
	folder.LineID++
	item := &FavItem{
		FavType: FavItemTypeLine,
		FavAttr: uint8(FavhFav),
		Item:    &FavLineItem{LineID: folder.LineID},
	}
	folder.FavItems = append(folder.FavItems, item)
	folder.NLines++
	folder.DataTail = uint16(len(folder.FavItems))
	folder.NAlloc = folder.DataTail + favPreAlloc
	return item, nil
}

func appendFavoriteFolder(folder *FavFolder, title string) (*FavItem, error) {
	if folder == nil {
		return nil, fmt.Errorf("pttbbs: favorite root folder is nil")
	}
	if folder.NFolders >= pttFavoriteMaxFolder {
		return nil, fmt.Errorf("pttbbs: favorite folder limit reached")
	}
	if title == "" {
		return nil, fmt.Errorf("pttbbs: favorite folder title is empty")
	}
	if len(utf8ToBig5UAOString(title)) > BoardTitleLength {
		return nil, fmt.Errorf("pttbbs: favorite folder title exceeds %d bytes in Big5", BoardTitleLength)
	}

	folder.FolderID++
	item := &FavItem{
		FavType: FavItemTypeFolder,
		FavAttr: uint8(FavhFav),
		Item: &FavFolderItem{
			FolderID: folder.FolderID,
			Title:    title,
			ThisFolder: &FavFolder{
				NAlloc:   favPreAlloc,
				FavItems: []*FavItem{},
			},
		},
	}
	folder.FavItems = append(folder.FavItems, item)
	folder.NFolders++
	folder.DataTail = uint16(len(folder.FavItems))
	folder.NAlloc = folder.DataTail + favPreAlloc
	return item, nil
}

func (c *Connector) appendFavoriteBoard(folder *FavFolder, boardID string) (*FavItem, error) {
	if folder == nil {
		return nil, fmt.Errorf("pttbbs: favorite root folder is nil")
	}
	if boardID == "" {
		return nil, fmt.Errorf("pttbbs: favorite board id is empty")
	}

	bid, canonicalID, err := c.lookupFavoriteBoard(boardID)
	if err != nil {
		return nil, err
	}
	for _, existing := range folder.FavItems {
		if board, ok := existing.Item.(*FavBoardItem); ok && board.BoardID == bid {
			return nil, fmt.Errorf("pttbbs: board %s is already in favorites", canonicalID)
		}
	}

	item := &FavItem{
		FavType: FavItemTypeBoard,
		FavAttr: uint8(FavhFav),
		Item: &FavBoardItem{
			BoardID:   bid,
			LastVisit: time.Unix(0, 0),
			boardID:   canonicalID,
		},
	}
	folder.FavItems = append(folder.FavItems, item)
	folder.NBoards++
	folder.DataTail = uint16(len(folder.FavItems))
	folder.NAlloc = folder.DataTail + favPreAlloc
	return item, nil
}

func (c *Connector) lookupFavoriteBoard(boardID string) (uint32, string, error) {
	path, err := c.GetBoardRecordsPath()
	if err != nil {
		return 0, "", fmt.Errorf("pttbbs: get board records path: %w", err)
	}
	data, err := os.ReadFile(path) // #nosec G304 -- path is the configured BBS .BRD path.
	if err != nil {
		return 0, "", fmt.Errorf("pttbbs: read board records: %w", err)
	}
	if len(data)%BoardHeaderRecordLength != 0 {
		return 0, "", fmt.Errorf("pttbbs: malformed board records size %d", len(data))
	}
	for offset := 0; offset < len(data); offset += BoardHeaderRecordLength {
		header, err := UnmarshalBoardHeader(data[offset : offset+BoardHeaderRecordLength])
		if err != nil {
			return 0, "", fmt.Errorf("pttbbs: decode board record: %w", err)
		}
		if strings.EqualFold(header.BoardID(), boardID) {
			return uint32(offset/BoardHeaderRecordLength + 1), header.BoardID(), nil
		}
	}
	return 0, "", fmt.Errorf("pttbbs: board %s not found", boardID)
}

func replaceFavoriteFile(path string, data []byte) error {
	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, ".fav.tmp-*")
	if err != nil {
		return fmt.Errorf("pttbbs: create favorite temp file: %w", err)
	}
	tmpPath := tmp.Name()
	defer os.Remove(tmpPath)

	if _, err := tmp.Write(data); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("pttbbs: write favorite temp file: %w", err)
	}
	if err := tmp.Sync(); err != nil {
		_ = tmp.Close()
		return fmt.Errorf("pttbbs: sync favorite temp file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		return fmt.Errorf("pttbbs: close favorite temp file: %w", err)
	}

	if err := os.Rename(tmpPath, path); err == nil {
		return nil
	}

	// Windows does not replace an existing destination with os.Rename. Keep a
	// rollback copy so the fallback never leaves the user's old favorites lost.
	backup, err := os.CreateTemp(dir, ".fav.backup-*")
	if err != nil {
		return fmt.Errorf("pttbbs: create favorite backup path: %w", err)
	}
	backupPath := backup.Name()
	if err := backup.Close(); err != nil {
		_ = os.Remove(backupPath)
		return fmt.Errorf("pttbbs: close favorite backup placeholder: %w", err)
	}
	if err := os.Remove(backupPath); err != nil {
		return fmt.Errorf("pttbbs: prepare favorite backup path: %w", err)
	}
	if err := os.Rename(path, backupPath); err != nil {
		return fmt.Errorf("pttbbs: backup existing favorite file: %w", err)
	}
	if err := os.Rename(tmpPath, path); err != nil {
		_ = os.Rename(backupPath, path)
		return fmt.Errorf("pttbbs: replace favorite file: %w", err)
	}
	_ = os.Remove(backupPath)
	return nil
}
