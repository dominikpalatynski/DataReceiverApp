package cache

import (
	"ConfigApp/config"
	"ConfigApp/logging"
	"ConfigApp/model"
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type Cache interface {
	SetDeviceDataToCache(model.DeviceData, string)
	GetDeviceDataFromCache(string) (*model.DeviceData, error)

	SetDeviceStateCredentialsToCache(model.DeviceStateCredentials, string)
	GetDeviceStateCredentialsToCache(string) (*model.DeviceStateCredentials, error)
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

	logging.Log.Info("Connected to Redis: %v", pong)

	return &RedisCache{
		client: redisClient,
	}, nil
}

func (c RedisCache) SetDeviceDataToCache(deviceData model.DeviceData, deviceKey string) {
	deviceJSON, err := json.Marshal(deviceData)
	if err != nil {
		fmt.Printf("Can not marshal in to json: %+v\n", err)
		return
	}

	ctx := context.Background()

	err = c.client.Set(ctx, deviceKey, deviceJSON, 0).Err()
	if err != nil {
		fmt.Printf("Can not set data in to cache: %+v\n", err)
		return
	}
	fmt.Print("DeviceData saved in cache \n")
}

func (c RedisCache) GetDeviceDataFromCache(deviceKey string) (*model.DeviceData, error) {

	ctx := context.Background()

	cachedData, err := c.client.Get(ctx, deviceKey).Result()

	if err != nil {
		return nil, errors.New("DeviceData not found")
	}

	var deviceData model.DeviceData

	if err := json.Unmarshal([]byte(cachedData), &deviceData); err != nil {
		return nil, err
	}

	fmt.Printf("Read data from cache: %+v\n", deviceData)

	return &deviceData, nil
}

func (c RedisCache) SetDeviceStateCredentialsToCache(deviceData model.DeviceStateCredentials, deviceKey string) {
	deviceJSON, err := json.Marshal(deviceData)
	if err != nil {
		fmt.Printf("Can not marshal in to json: %+v\n", err)
		return
	}

	ctx := context.Background()

	err = c.client.Set(ctx, deviceKey, deviceJSON, 0).Err()
	if err != nil {
		fmt.Printf("Can not set data in to cache: %+v\n", err)
		return
	}
	fmt.Print("DeviceStateCredentials saved in cache \n")
}

func (c RedisCache) GetDeviceStateCredentialsToCache(deviceKey string) (*model.DeviceStateCredentials, error) {

	ctx := context.Background()

	cachedData, err := c.client.Get(ctx, deviceKey).Result()

	if err != nil {
		return nil, errors.New("DeviceStateCredentials not found")
	}

	var deviceData model.DeviceStateCredentials

	if err := json.Unmarshal([]byte(cachedData), &deviceData); err != nil {
		return nil, err
	}

	fmt.Printf("Read data from cache: %+v\n", deviceData)

	return &deviceData, nil
}