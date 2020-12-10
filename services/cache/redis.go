package cache

import (
	"log"
	"strconv"

	"go-easy-frame/config"

	"github.com/go-redis/redis"
)

const Nil = redis.Nil

var (
	RedisClient *redis.Client
)

func init() {
	InitCache()
}

func New() *redis.Client {
	return RedisClient
}

func InitCache() {
	InitCliRedis()
}

func InitCliRedis() {
	log.Println("init InitCliRedis")
	cfg := config.GetEnv()

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + strconv.Itoa(cfg.Redis.Port),
		Password: cfg.Redis.Pwd,
		DB:       cfg.Redis.Db,
	})

	_, err := RedisClient.Ping().Result()
	if nil != err {
		log.Println("初始化redis失败", err)
		panic(err)
	}
}
