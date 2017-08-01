package redis

import (
	"fmt"
	"github.com/go-redis/redis"
)

var connection *redis.Client

func InitConnection() *redis.Client {
	connection = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return connection
}

func GetConnection() (*redis.Client, error) {
	var err error
	if connection == nil {
		err = fmt.Errorf("No created connection with Redis")
	}
	return connection, err
}
