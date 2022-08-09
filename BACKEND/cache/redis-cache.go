package cache

import (
	"encoding/json"
	"learning-go/entity"
	"time"

	"github.com/go-redis/redis/v7"
)

type UrlCache interface {
	Set(key string, value entity.FreeUrl)
	Get(key string) *entity.FreeUrl
}

type urlCache struct {
	host    string
	db      int
	expires time.Duration
}

func NewUrlCache(host string, db int, expires time.Duration) UrlCache {
	return &urlCache{
		host:    host,
		db:      db,
		expires: expires,
	}
}

func (cache *urlCache) getClient()  *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cache.host,
		//Password: "q",
		DB: cache.db,
	})
}

func (cache *urlCache) Set(key string, value entity.FreeUrl) {
	client := cache.getClient()

	json, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	client.Set(key, string(json), cache.expires*time.Second)
}

func (cache *urlCache) Get(key string) *entity.FreeUrl {

	client := cache.getClient()
	val, err := client.Get(key).Result()
	if err != nil {
		return nil
	}
	url := entity.FreeUrl{}
	err = json.Unmarshal([]byte(val), &url)
	if err != nil {
		panic(err)
	}
	return &url
}
