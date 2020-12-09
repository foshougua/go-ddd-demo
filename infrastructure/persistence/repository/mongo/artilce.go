package mongo

import (
	"ddd-test/domain/entity"
	"ddd-test/domain/repository"
	"ddd-test/infrastructure/config"
	"ddd-test/infrastructure/persistence/translator"
	"gopkg.in/mgo.v2"
)

type ArticleMongoRepo struct {
	mgo *mgo.Session
}

func NewArticleMongoRepo(session *mgo.Session) *ArticleMongoRepo {
	return &ArticleMongoRepo{
		mgo: session,
	}
}

var _ repository.ArticleRepository = &ArticleMongoRepo{}

func (amr *ArticleMongoRepo) CreateArticle(info *entity.Article) error {
	tr := translator.NewMongoTranslator()
	mgoInfo := tr.CovertArticleToMongo(info)
	return amr.mgo.DB(config.Cfg.DBName).C("article").Insert(mgoInfo)
}

func (amr *ArticleMongoRepo) DeleteArticle(id string) error {
	return nil
}

func (amr *ArticleMongoRepo) UpdateArticle(upInfo map[string]interface{}) error {
	return nil
}

func (amr *ArticleMongoRepo) GetArticleInfoById(id string) (*entity.Article, error) {
	return nil, nil
}

func (amr *ArticleMongoRepo) GetArticleList(pageNo, pageSize int, searchText string) ([]entity.Article, int, error) {
	return nil, 0, nil
}
