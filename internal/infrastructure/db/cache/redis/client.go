package redis

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	redisClient *Client
	redisOnce   sync.Once
)

// Client wraps Redis client with additional functionality
type Client struct {
	*redis.Client
	ctx context.Context
}

// NewRedisClient creates a singleton Redis client
func NewRedisClient(address, password string, db int) *Client {
	redisOnce.Do(func() {
		client := redis.NewClient(&redis.Options{
			Addr:     address,
			Password: password,
			DB:       db,
		})

		ctx := context.Background()
		if err := client.Ping(ctx).Err(); err != nil {
			fmt.Printf("❌ Cannot connect to Redis: %v\n", err)
			panic(fmt.Errorf("redis connection failed: %w", err))
		}

		fmt.Println("✅ Successfully connected to Redis")
		redisClient = &Client{
			Client: client,
			ctx:    ctx,
		}
	})
	return redisClient
}

// Set stores a value with expiration
func (c *Client) Set(key string, value interface{}, expiration time.Duration) error {
	return c.Client.Set(c.ctx, key, value, expiration).Err()
}

// Get retrieves a value by key
func (c *Client) Get(key string) (string, error) {
	return c.Client.Get(c.ctx, key).Result()
}

// Del deletes keys
func (c *Client) Del(keys ...string) error {
	return c.Client.Del(c.ctx, keys...).Err()
}

// Exists checks if key exists
func (c *Client) Exists(key string) (bool, error) {
	result, err := c.Client.Exists(c.ctx, key).Result()
	return result > 0, err
}

// Ping tests the connection
func (c *Client) Ping() error {
	return c.Client.Ping(c.ctx).Err()
}

// Close closes the Redis connection
func (c *Client) Close() error {
	return c.Client.Close()
}
