package cache

import (
	"carrick-js-api/pkgs/config"
	"carrick-js-api/pkgs/logger"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"sync"
	"time"
)

type RedisCacheSingleton interface {
	GetCache() int
}

type redisCacheSingleton struct {
	client *redis.Client
}

var (
	redisCacheOnce     sync.Once
	redisCacheInstance *redisCacheSingleton
)

func GetRedisCacheInstance() *redisCacheSingleton {
	logger := logger.GetLoggerInstance()

	if redisCacheInstance == nil {
		redisCacheOnce.Do(
			func() {
				redisCacheInstance = &redisCacheSingleton{}

				client := redis.NewClient(&redis.Options{
					Addr:     fmt.Sprintf("%s:%s", config.AppConfig.Redis.Host, config.AppConfig.Redis.Port),
					Password: config.AppConfig.Redis.Password, // no password set
					DB:       config.AppConfig.Redis.DB,       // use default DB
				})

				_, err := client.Ping().Result()

				if err == nil {
					redisCacheInstance.client = client
				} else {
					logger.Error(err)
				}
			})
	}

	return redisCacheInstance
}

func (c *redisCacheSingleton) GetClient() *redis.Client {
	return c.client
}

func (c *redisCacheSingleton) Set(key string, value interface{}, ttl int) error {
	ttlCache := time.Duration(ttl) * time.Second

	if c.client == nil {
		return errors.New("Cache client not initialized")
	}

	_, err := c.client.Ping().Result()
	if err != nil {
		return err
	}

	valueJson, err := json.Marshal(value)
	if err != nil {
		return err
	}

	prefix := config.AppConfig.Redis.Prefix
	err = c.client.Set(prefix+key, valueJson, ttlCache).Err()
	if err != nil {
		return err
	} else {
		return nil
	}
}

func (c *redisCacheSingleton) Get(key string, dest interface{}) error {
	if c.client == nil {
		return errors.New("Cache client not initialized")
	}

	_, err := c.client.Ping().Result()
	if err != nil {
		return err
	}

	prefix := config.AppConfig.Redis.Prefix
	value, err := c.client.Get(prefix+key).Result()
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(value), dest)
}