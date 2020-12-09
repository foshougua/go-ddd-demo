package handler

import (
	"ddd-test/application/service"
	"ddd-test/domain/entity"
	"ddd-test/infrastructure/utility"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ArticleHandlers struct {
	articleAppSvr service.ArticleAppInterface
}

func NewArticleHandlers() *ArticleHandlers {
	return &ArticleHandlers{}
}

func (a *ArticleHandlers) Create(c *gin.Context) {
	req := entity.CreateArticleReq{}
	errRsp := make(map[string]interface{})
	if err := c.ShouldBindJSON(&req); err != nil {
		utility.MakeLocRsp(c, http.StatusBadRequest, utility.SERVER_ERR, errRsp)
		return
	}
	info := &entity.Article{
		Author:     req.Author,
		Content:    req.Content,
		CreateTime: time.Now().UnixNano() / 1e6,
	}
	svr := service.NewArticleSvr()
	svrRsp := svr.Create(c, info)
	svrRsp.MakeRsp(c, errRsp)
	return
}

func (a *ArticleHandlers) Delete(c *gin.Context) {
	return
}

func (a *ArticleHandlers) Update(c *gin.Context) {
	return
}

func (a *ArticleHandlers) GetInfoById(c *gin.Context) {
	errRsp := make(map[string]interface{})
	id := c.Param("id")
	if id == "" {
		utility.MakeLocRsp(c, http.StatusBadRequest, utility.SERVER_ERR, errRsp)
		return
	}
	data, svrRsp := a.articleAppSvr.GetInfoById(c, id)
	if svrRsp.ErrCode != utility.OK {
		svrRsp.MakeRsp(c, errRsp)
		return
	}
	//也可做一层dto
	svrRsp.MakeRsp(c, data)
	return
}

func (a *ArticleHandlers) GetList(c *gin.Context) {
	return
}
