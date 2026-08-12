package bbs

import "fmt"

// DeleteArticleConnector is implemented by drivers that support deleting board
// articles while preserving the deletion semantics of the underlying BBS.
type DeleteArticleConnector interface {
	// DeleteBoardArticle deletes filename from boardID. deletedBy identifies the
	// user performing the deletion so drivers can preserve native audit/title
	// semantics when applicable.
	DeleteBoardArticle(boardID, filename, deletedBy string) error
}

// DeleteBoardArticle deletes an article using the native deletion semantics of
// the configured connector.
func (db *DB) DeleteBoardArticle(boardID, filename, deletedBy string) error {
	c, ok := db.connector.(DeleteArticleConnector)
	if !ok {
		return fmt.Errorf("bbs: connector don't support DeleteArticleConnector")
	}
	return c.DeleteBoardArticle(boardID, filename, deletedBy)
}
