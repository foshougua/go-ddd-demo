package main

import (
	"ddd-test/infrastructure/config"
	"ddd-test/infrastructure/logger"
	"ddd-test/infrastructure/persistence/repository/mongo"
	"ddd-test/infrastructure/persistence/repository/sql"
	"ddd-test/infrastructure/redis"
	"ddd-test/interfaces"
	"github.com/gin-gonic/gin"
)

func main() {
	mongo.StartMongo(config.Cfg.MongoURL)
	sql.StartUp(config.Cfg.SqlAddr)
	redis.StartRedis()
	gin.SetMode("debug")
	eng := gin.Default()
	interfaces.RouterLoad(eng)
	logger.GLog.Fatalln(eng.Run(":8080"))
}
