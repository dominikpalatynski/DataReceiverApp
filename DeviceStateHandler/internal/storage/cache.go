package storage

import (
	"DeviceStateHandler/internal/model"

	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
    SetDeviceState(model.DeviceState, string) error
    GetDeviceState(string) (*model.DeviceState, error)
}

type RedisCache struct {
    client *redis.Client
}

func NewRedisClient(redisUrl string, redisPassword string) (*RedisCache, error) {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisUrl, 
		Password: redisPassword,
		DB:       0,
	})
	  
	ctx := context.Background()
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return nil, err
	}

	log.Printf("Connected to Redis: %v", pong)

	return &RedisCache{
		client: redisClient,
	}, nil
}

func (r *RedisCache) SetDeviceState(state model.DeviceState, key string) error{
    ctx := context.Background()
    data, err := json.Marshal(state)
    if err != nil {
        return err
    }
    r.client.Set(ctx, key, data, 24*time.Hour)
	log.Printf("Device state: %s saved in cache", state.DeviceState)
	return nil
}

func (r *RedisCache) GetDeviceState(key string) (*model.DeviceState, error) {
    ctx := context.Background()
    data, err := r.client.Get(ctx, key).Result()
    if err == redis.Nil {
        return nil, nil
    } else if err != nil {
        return nil, err
    }

    var state model.DeviceState
    if err := json.Unmarshal([]byte(data), &state); err != nil {
        return nil, err
    }

    return &state, nil
}
