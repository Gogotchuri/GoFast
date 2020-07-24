package cache

import (
	"strconv"
	"sync"

	"github.com/Gogotchuri/GoFast/config"
	"github.com/go-redis/redis"
)

var client *redis.Client
var once sync.Once

/*GetRedisInstance Returns redis singleton client*/
func GetRedisInstance() *redis.Client {
	once.Do(func() {
		client = newClient(config.GetInstance())
	})
	return client
}

/*GetRedisInstance Creates and returns redis singleton client*/
func newClient(cfg *config.Config) *redis.Client {
	//Initializing redis
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + strconv.Itoa(cfg.Redis.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
