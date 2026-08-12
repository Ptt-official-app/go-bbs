package bbs

import "fmt"

// FavoriteCreateOptions describes one favorite item to append to the user's
// root favorite folder.
type FavoriteCreateOptions struct {
	Type    FavoriteType
	BoardID string
	Title   string
}

// WriteFavoriteConnector is implemented by drivers that can persist user
// favorite items using the native BBS file format.
type WriteFavoriteConnector interface {
	AddUserFavorite(userID string, options FavoriteCreateOptions) (FavoriteRecord, error)
}

// AddUserFavorite appends one favorite item using the configured connector.
func (db *DB) AddUserFavorite(userID string, options FavoriteCreateOptions) (FavoriteRecord, error) {
	connector, ok := db.connector.(WriteFavoriteConnector)
	if !ok {
		return nil, fmt.Errorf("bbs: connector don't support WriteFavoriteConnector")
	}
	return connector.AddUserFavorite(userID, options)
}
