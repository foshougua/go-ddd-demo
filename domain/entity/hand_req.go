package entity

type CreateArticleReq struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}
