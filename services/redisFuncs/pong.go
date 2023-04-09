package redisfuncs

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

func Pong(ctx context.Context, c *redis.Client) {
	pong, err := c.Ping(ctx).Result()

	if err != nil {
		fmt.Println(err)
	}

	log.Println(pong)
}
