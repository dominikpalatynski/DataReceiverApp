package cache

import (
	"ConfigApp/config"
	"ConfigApp/model"
	"fmt"

	"context"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	SetDeviceDataToCache(model.DeviceData) error
	GetDeviceDataFromCache(int) (model.DeviceData, error)
}

type RedisCache struct {
	client *redis.Client
}

func NewRedisClient(config config.Config) (*RedisCache, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.Cache.Url, 
		Password: config.Cache.Password,
		DB:       0,
	   })
	  
	   ctx := context.Background()
	   pong, err := redisClient.Ping(ctx).Result()
	   if err != nil {
			return nil, err
	   }

	   fmt.Println("Connected to Redis:", pong)

	return &RedisCache{
		client: redisClient,
	}, nil
}

func (c RedisCache) SetDeviceDataToCache(model.DeviceData) error {
	return nil
}

func (c RedisCache) GetDeviceDataFromCache(int) (model.DeviceData, error) {
	return model.DeviceData{}, nil
}
