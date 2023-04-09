package redisfuncs

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func Set(ctx context.Context, client *redis.Client, key, value string, expiration time.Duration) {
	err := client.Set(ctx, key, value, expiration).Err()
	// if there has been an error setting the value
	// handle the error
	if err != nil {
		fmt.Println(err)
	}
}
