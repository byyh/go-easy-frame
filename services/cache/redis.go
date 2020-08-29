package cache

import (
	"fmt"
	"log"
	"strconv"

	"github.com/byyh/go/com"
	"github.com/go-redis/redis"
	"go-easy-frame/config"
)

const Nil = redis.Nil

var (
	RedisClient *redis.Client
)

func New() *redis.Client {
	return RedisClient
}

func InitCache() {
	InitCliRedis()
}

func InitCliRedis() {
	log.Println("init InitCliRedis")
	cfg := config.GetEnv()
	dbNum, err := strconv.ParseInt(cfg.Redis.Db), 10, 32)
	com.CheckErr(err)

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host + ":" + strconv.Itoa(cfg.Redis.Port),
		Password: cfg.Redis.Pwd,
		DB:       int(dbNum),
	})

	_, err = RedisClient.Ping().Result()
	if nil != err {
		log.Println("初始化redis失败", err)
		panic(err)
	}
}
