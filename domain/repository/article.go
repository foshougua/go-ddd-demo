package repository

import "ddd-test/domain/entity"

type ArticleRepository interface {
	CreateArticle(info *entity.Article) error
	DeleteArticle(id string) error
	UpdateArticle(upInfo map[string]interface{}) error
	GetArticleInfoById(id string) (*entity.Article, error)
	GetArticleList(pageNo, pageSize int, searchText string) ([]entity.Article, int, error)
}
