package redisfuncs

import (
	"context"
	"log"
	"os"

	r "github.com/go-redis/redis/v8"
)

func Connect() (client *r.Client) {
	client = r.NewClient(&r.Options{
		Addr:     os.Getenv("REDIS"),
		Password: "",
		DB:       0,
	})

	if ping := client.Ping(context.TODO()); ping.Err() != nil {
		log.Panic("Redis is not available: ", ping.Err().Error())
	}

	return
}
