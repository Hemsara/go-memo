package initializers

import (
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var RD *redis.Client

func InitRedis() {
	redisScheme := os.Getenv("REDIS_SCHEME")
	redisUsername := os.Getenv("REDIS_USERNAME")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	redisURL := fmt.Sprintf("%s://%s:%s@%s:%s", redisScheme, redisUsername, redisPassword, redisHost, redisPort)

	opt, _ := redis.ParseURL(redisURL)
	c := redis.NewClient(opt)

	RD = c

}
