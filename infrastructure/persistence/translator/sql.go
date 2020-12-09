package translator

import (
	"ddd-test/domain/entity"
	"ddd-test/infrastructure/persistence/entity/sql"
)

type SqlTranslatorInterface interface {
	CovertArticleToEntity(info *sql.ArticleInfo) *entity.Article
	CovertArticleToSql(info *entity.Article) *sql.ArticleInfo
}

type SqlTranslator struct{}

func NewSqlTranslator() *SqlTranslator {
	return &SqlTranslator{}
}

var _ SqlTranslatorInterface = &SqlTranslator{}

func (st *SqlTranslator) CovertArticleToEntity(info *sql.ArticleInfo) *entity.Article {
	return nil
}

func (st *SqlTranslator) CovertArticleToSql(info *entity.Article) *sql.ArticleInfo {
	return nil
}
