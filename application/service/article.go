package service

import (
	"ddd-test/domain/entity"
	domainSvr "ddd-test/domain/service"
	"ddd-test/infrastructure/utility"
	"github.com/gin-gonic/gin"
)

type ArticleAppInterface interface {
	Create(c *gin.Context, info *entity.Article) *utility.SvrRspInfo
	Delete(c *gin.Context)
	Update(c *gin.Context)
	GetInfoById(c *gin.Context, id string) (*entity.Article, *utility.SvrRspInfo)
	GetList(c *gin.Context)
}

type ArticleApp struct {
	svr domainSvr.ArticleServiceInterface
}

//用做类型断言，如果ArticleSvr未实现接口则报错
var _ ArticleAppInterface = &ArticleApp{}

func NewArticleSvr() *ArticleApp {
	return &ArticleApp{}
}

func (a *ArticleApp) Create(c *gin.Context, info *entity.Article) *utility.SvrRspInfo {
	return a.svr.CreateArticle(info)
}

func (a *ArticleApp) Delete(c *gin.Context) {
}

func (a *ArticleApp) Update(c *gin.Context) {
}

func (a *ArticleApp) GetInfoById(c *gin.Context, id string) (*entity.Article, *utility.SvrRspInfo) {
	return a.svr.GetArticle(id)
}

func (a *ArticleApp) GetList(c *gin.Context) {
}
