package services

import (
	"context"
	"encoding/json"
	funcs "klineService/services/redisFuncs"
	"log"
	"strconv"

	"github.com/go-redis/redis/v8"
)

func (r *Redis) InitCache() {
	CacheInstance.Client = funcs.Connect()

	conn := r.Database.CreateConnection()
	assets := r.AssetRepo.List(nil, conn)
	conn.Close()

	b, err := json.Marshal(assets)

	if err != nil {
		log.Panic("redis marshal: ", err)
	}

	CacheInstance.Client.Set(context.TODO(), "Assets", b, 0)
}

func (r *Redis) GetInstance() *redis.Client {
	return CacheInstance.Client
}

func (r *Redis) GetCache(key, primitiveType string) (val any) {
	cachedVal, err := CacheInstance.Client.Get(context.TODO(), key).Result()

	if err != nil {
		// log.Println("Redis fetch cache error: ", err)
	}

	switch primitiveType {
	case "bool":
		if cachedVal == "" {
			return
		}

		val, err = strconv.ParseBool(cachedVal)

		if err != nil {
			log.Panic("Redis GetCache Conv: ", err)
		}
	case "float64":
		if cachedVal == "" {
			return
		}

		val, err = strconv.ParseFloat(cachedVal, 64)

		if err != nil {
			log.Panic("Redis GetCache Conv: ", err)
		}
	case "int64":
		if cachedVal == "" {
			return
		}

		val, err = strconv.ParseInt(cachedVal, 10, 64)

		if err != nil {
			log.Panic("Redis GetCache Conv: ", err)
		}
	case "string":
		if cachedVal == "" {
			return
		}

		return cachedVal
	}

	return val
}
