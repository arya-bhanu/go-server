package database

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"

	"go-server/utils"
)

var Redis *redis.Client
var ctx = context.Background()

func init() {
	Redis = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	// Ping Redis to test connection
	pong, err := Redis.Ping(ctx).Result()
	if err != nil {
		utils.Loger.Error(fmt.Sprintf("❌ Redis not connected:%s", err.Error()))
		return
	}

	utils.Loger.Info(fmt.Sprintf("✅ Redis connected:%s", pong))
}
