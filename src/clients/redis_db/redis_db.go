package redis_db

import 	"github.com/go-redis/redis/v8"

var DB *redis.Client

func init() {
	DB = redis.NewClient(&redis.Options{
		Addr:     "172.17.0.2:6379",
		Password: "",
		DB:       0,
	},
	)
}
