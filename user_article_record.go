package bbs

type UserArticleRecord interface {
	BoardID() string
	Title() string
	Owner() string
	ArticleID() string
}

type userArticleRecord map[string]string

func (r userArticleRecord) BoardID() string {
	return r["board_id"]
}
func (r userArticleRecord) Title() string {
	return r["title"]
}
func (r userArticleRecord) Owner() string {
	return r["owner"]
}
func (r userArticleRecord) ArticleID() string {
	return r["article_id"]
}
