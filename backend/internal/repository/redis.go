package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"backend/internal/models"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStorage(addr string, password string, db int) *RedisStorage {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	return &RedisStorage{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (r *RedisStorage) Save(stat models.Stat) error {
	data, err := json.Marshal(stat)
	if err != nil {
		return fmt.Errorf("failed to marshal data stat: %w", err)
	}

	key := fmt.Sprintf("stat:%s%d", stat.Symbol, stat.Timedump)

	err = r.client.Set(r.ctx, key, data, 24*time.Hour).Err()
	if err != nil {
		return fmt.Errorf("failed to save to redis: %w", err)
	}

	return nil
}
