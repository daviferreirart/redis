package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func main() {

	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "SUASENHA",
		DB:       0,
	})

	ctx := context.Background()

	err := redisClient.Set(ctx, "name", "davi", 0).Err()
	if err != nil {
		fmt.Print(err)
	}

	_, err = redisClient.Get(ctx, "name").Result()
	if err != nil {
		fmt.Print(err)
	}

	list := map[string]string{"nome": "john", "sobrenome": "wick", "cidade": "new york", "uuid": uuid.New().String()}

	for k, v := range list {
		err = redisClient.HSet(ctx, "values", k, v).Err()
		if err != nil {
			fmt.Print(err)
		}
	}
	fmt.Println(redisClient.HGetAll(ctx, "values").Val())
	redisClient.HDel(ctx, "values", "cidade", "uuid")
	fmt.Println(redisClient.HGetAll(ctx, "values").Val())

}
