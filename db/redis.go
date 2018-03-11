package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

var (
	gameClient *redis.Client
)

func init() {
	gameClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := gameClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Initialize redis error:%v", err))
	}
}

func SetGame(key string, value interface{}) error {
	return gameClient.Set(key, value, 0).Err()
}

func GetGame(key string) (string, error) {
	return gameClient.Get(key).Result()
}
