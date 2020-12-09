package cache

import (
	"ddd-test/domain/entity"
	"ddd-test/infrastructure/redis"
	"encoding/json"
)

type ArticleCacheInterface interface {
	Set(info *entity.Article) error
	Get(id string) (*entity.Article, error)
}

type ArticleCache struct {
	ro redis.RedisOperInterfaces
}

func NewArticleCache(r redis.RedisOperInterfaces) *ArticleCache {
	return &ArticleCache{
		ro: r,
	}
}

func (a *ArticleCache) Set(info *entity.Article) error {
	value, _ := json.Marshal(info)
	return a.ro.Set(string(value))
}

func (a *ArticleCache) Get(id string) (*entity.Article, error) {
	key := a.genKey(id)
	value, err := a.ro.Get(key)
	if err != nil {
		return nil, err
	}
	info := &entity.Article{}
	err = json.Unmarshal([]byte(value), info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

func (a *ArticleCache) genKey(id string) string {
	return id
}
