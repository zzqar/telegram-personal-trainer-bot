package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"telegram-personal-trainer-bot/internal/config"
)

type StateManager struct {
	client *redis.Client
	ctx    context.Context
}

func MustRedis(conf config.Redis) *StateManager {
	rdb := redis.NewClient(&redis.Options{
		Addr:     conf.Addr,
		Password: conf.Password,
		DB:       conf.DB,
	})

	ctx := context.Background()
	return &StateManager{client: rdb, ctx: ctx}
}

// Set
// Установить состояние пользователя
func (r *StateManager) Set(key, state string) error {
	return r.client.Set(r.ctx, key, state, 0).Err()
}

// Get
// Получить состояние пользователя
func (r *StateManager) Get(key string) (string, error) {
	return r.client.Get(r.ctx, key).Result()
}

//// GetInt
//// Получить состояние пользователя
//func (r *StateManager) GetInt(key string) (int, error) {
//	return r.Get(key).Int()
//}
