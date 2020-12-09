package sql

import (
	"ddd-test/domain/entity"
	"ddd-test/domain/repository"
	"ddd-test/infrastructure/persistence/translator"
	"github.com/jinzhu/gorm"
)

type ArticleSqlRepo struct {
	db *gorm.DB
}

func NewArticleSqlRepo(db *gorm.DB) *ArticleSqlRepo {
	return &ArticleSqlRepo{
		db: db,
	}
}

var _ repository.ArticleRepository = &ArticleSqlRepo{}

func (amr *ArticleSqlRepo) CreateArticle(info *entity.Article) error {
	tr := translator.NewSqlTranslator()
	sqlInfo := tr.CovertArticleToSql(info)
	err := amr.db.Create(sqlInfo).Error
	return err
}

func (amr *ArticleSqlRepo) DeleteArticle(id string) error {
	return nil
}

func (amr *ArticleSqlRepo) UpdateArticle(upInfo map[string]interface{}) error {
	return nil
}

func (amr *ArticleSqlRepo) GetArticleInfoById(id string) (*entity.Article, error) {
	return nil, nil
}

func (amr *ArticleSqlRepo) GetArticleList(pageNo, pageSize int, searchText string) ([]entity.Article, int, error) {
	return nil, 0, nil
}
