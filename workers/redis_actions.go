package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func InitRedis(addr string) {
	client = redis.NewClient(&redis.Options{
		Addr: addr, // Redis server address
	})
	ctx := context.Background()
	if _, err := client.Ping(ctx).Result(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
}

// AddToRedis adds a new key with a value and expiration to the Redis cache.
func AddToRedis(key, value string, expiration time.Duration) {
	ctx := context.Background()
	if _, err := client.Set(ctx, key, value, expiration).Result(); err != nil {
		log.Printf("Failed to add key %s to Redis: %v", key, err)
	}
}

// PrintRedisCache prints all keys and their values stored in the Redis cache.
func PrintRedisCache() {
	ctx := context.Background()
	keys, err := client.Keys(ctx, "*").Result()
	if err != nil {
		log.Printf("Failed to retrieve keys from Redis: %v", err)
		return
	}
	for _, key := range keys {
		value, err := client.Get(ctx, key).Result()
		if err != nil {
			log.Printf("Failed to retrieve value for key %s: %v", key, err)
			continue
		}
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}
}

// ClearRedisCache clears all keys in the Redis cache.
func ClearRedisCache() {
	ctx := context.Background()
	keys, err := client.Keys(ctx, "*").Result()
	if err != nil {
		log.Printf("Failed to retrieve keys for clearing Redis: %v", err)
		return
	}
	for _, key := range keys {
		if _, err := client.Del(ctx, key).Result(); err != nil {
			log.Printf("Failed to delete key %s from Redis: %v", key, err)
		}
	}
}

// RefreshRedisKey refreshes the expiration time of a key in Redis.
func RefreshRedisKey(key string, expiration time.Duration) {
	ctx := context.Background()
	if _, err := client.Expire(ctx, key, expiration).Result(); err != nil {
		log.Printf("Failed to refresh expiration for key %s in Redis: %v", key, err)
	}
}
