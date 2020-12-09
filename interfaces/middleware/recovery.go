package middleware

import (
	"ddd-test/infrastructure/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Recovery(c *gin.Context) {
	logger.GLog.Errorf("server panic ...")
	defer func() {
		if err := recover(); err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"errcode": "SERVER_INTERNAL_ERROR",
				"errmsg":  "服务内部错误",
				"data":    make(map[string]interface{}),
			})
		}
	}()
}
