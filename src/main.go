package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "SUASENHA",
		DB:       0,
	})

	err := redisClient.Set("nome", "davi", 0).Err()
	if err != nil {
		fmt.Print(err)
	}

	value, err := redisClient.Get("nome").Result()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(value)

}
