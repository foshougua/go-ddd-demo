package translator

import (
	"ddd-test/domain/entity"
	"ddd-test/infrastructure/persistence/entity/mongo"
)

type MongoTranslatorInterface interface {
	CovertArticleToEntity(info *mongo.ArticleInfo) *entity.Article
	CovertArticleToMongo(info *entity.Article) *mongo.ArticleInfo
}

type MongoTranslator struct {
}

var _ MongoTranslatorInterface = &MongoTranslator{}

func NewMongoTranslator() *MongoTranslator {
	return &MongoTranslator{}
}

func (mt *MongoTranslator) CovertArticleToEntity(info *mongo.ArticleInfo) *entity.Article {
	return nil
}

func (mt *MongoTranslator) CovertArticleToMongo(info *entity.Article) *mongo.ArticleInfo {
	return nil
}
