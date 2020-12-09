package service

import (
	"ddd-test/domain/entity"
	"ddd-test/domain/repository"
	"ddd-test/infrastructure/cache"
	"ddd-test/infrastructure/config"
	"ddd-test/infrastructure/logger"
	"ddd-test/infrastructure/persistence/repository/mongo"
	"ddd-test/infrastructure/persistence/repository/sql"
	"ddd-test/infrastructure/utility"
	"net/http"
)

type ArticleServiceInterface interface {
	CreateArticle(info *entity.Article) *utility.SvrRspInfo
	GetArticle(articleId string) (*entity.Article, *utility.SvrRspInfo)
}

type ArticleService struct {
	ar repository.ArticleRepository
	ac cache.ArticleCacheInterface
}

var _ ArticleServiceInterface = &ArticleService{}

func NewArticleService(ar repository.ArticleRepository, ac cache.ArticleCacheInterface) *ArticleService {
	s := &ArticleService{}
	switch config.Cfg.DBMode {
	case "mysql":
		s.ar = sql.NewArticleSqlRepo(sql.DB)
	default:
		s.ar = mongo.NewArticleMongoRepo(mongo.MgoDB)
	}
	s.ac = ac
	return s
}

func (as *ArticleService) CreateArticle(info *entity.Article) *utility.SvrRspInfo {
	err := as.ar.CreateArticle(info)
	if err != nil {
		logger.GLog.Errorf("create article err:", err.Error())
		return utility.NewLocSvrRspInfo(http.StatusInternalServerError, utility.SERVER_ERR)
	}
	return utility.NewLocSvrRspInfo(http.StatusOK, utility.OK)
}

func (as *ArticleService) GetArticle(articleId string) (*entity.Article, *utility.SvrRspInfo) {
	info, err := as.ac.Get(articleId)
	if err == nil {
		return info, utility.NewLocSvrRspInfo(http.StatusOK, utility.OK)
	}
	info, err = as.ar.GetArticleInfoById(articleId)
	if err != nil {
		logger.GLog.Errorf("get article from db err:", err.Error())
		return nil, utility.NewLocSvrRspInfo(http.StatusInternalServerError, utility.SERVER_ERR)
	}
	return info, utility.NewLocSvrRspInfo(http.StatusOK, utility.OK)
}
