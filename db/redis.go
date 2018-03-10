package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

var (
	roomClient *redis.Client
	gameClient *redis.Client
)

func init() {
	roomClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := roomClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Initialize redis error:%v", err))
	}

	gameClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	_, err = gameClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("Initialize redis error:%v", err))
	}
}

func SetRoom(key string, value interface{}) error {
	return roomClient.Set(key, value, 0).Err()
}

func GetRoom(key string) (string, error) {
	return roomClient.Get(key).Result()
}

func GetRooms() (string, error) {
	Get := func(client *redis.Client) *redis.StringCmd {
		cmd := redis.NewStringCmd("key *")
		client.Process(cmd)
		return cmd
	}
	return Get(roomClient).Result()
}

func SetGame(key string, value interface{}) error {
	return gameClient.Set(key, value, 0).Err()
}

func GetGame(key string) (string, error) {
	return gameClient.Get(key).Result()
}
