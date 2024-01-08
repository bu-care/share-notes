package get_key

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"testing"
)

func TestGet(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	res, err := client.Get(context.Background(), "key").Result()
	if err == redis.Nil {
		fmt.Println("111111")
		fmt.Println("err: ", err)
	}
	fmt.Printf("res: %p", &res)

}
