package interfaces

import (
	"ddd-test/interfaces/handler"
	"ddd-test/interfaces/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RouterLoad(g *gin.Engine) *gin.Engine {
	// Middlewares
	g.Use(middleware.Recovery)

	// 404 Handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	// 注意使用rest风格url
	article := handler.NewArticleHandlers()
	articleGroup := g.Group("/test-article")
	{
		articleGroup.POST("/info", article.Create)
		articleGroup.DELETE("/:id", article.Delete)
		articleGroup.PUT("/:id", article.Update)
		articleGroup.GET("/:id", article.GetInfoById)
		articleGroup.GET("/list", article.GetList)
	}

	return g
}
