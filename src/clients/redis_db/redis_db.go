package redis_db

import 	"github.com/go-redis/redis/v8"

var DB *redis.Client

func init() {
	DB = redis.NewClient(&redis.Options{
		Addr:     "localhost:8282",
		Password: "",
		DB:       0,
	},
	)
}
