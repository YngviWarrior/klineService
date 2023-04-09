package redisfuncs

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

func Get(ctx context.Context, client *redis.Client, key string) (val string) {
	val, err := client.Get(ctx, key).Result()

	if err != nil {
		fmt.Println(err)
	}

	return
}
