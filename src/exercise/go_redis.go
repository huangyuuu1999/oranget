package main

import (
	"context"
	"fmt"
	// "time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func Example() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})
	err := rdb.Set(ctx, "name", "June", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis get result:", val)

	val2, err := rdb.Get(ctx, "age").Result()
	if err == redis.Nil {
		fmt.Println("key age doesn't exist.")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("age", val2)
	}

	s1 := []string{"hello", "redis", "by", "go-redis", "Lib."}
	num, err := rdb.LPush(ctx, "tasks", s1).Result()
	fmt.Printf("tasks contains %d task.\n", num)

}

func main() {
	Example()
}