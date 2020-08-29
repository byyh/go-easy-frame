package config

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"

	"github.com/timest/env"
)

var (
	cfg *Config
)

// 需要根据实际情况配置该结构图属性
type Config struct {
	App     string
	Port    int      `default:"8000"`
	IsDebug bool     `env:"DEBUG"`
	Hosts   []string `slice_sep:","`
	Timeout int64    `default:"10"`

	Redis struct {
		Host string
		Port int
		Db   int
		Pwd  string
	}

	MysqlDns string

	// 队列，多个可分别定义
	RabbitmqPush struct {
		AmqpUri      string
		Queuename    string
		Exchange     string
		ExchangeType string
		RoutingKey   string `default:"10"`
	}

	// 日志系统
	Isfluentlog bool
	Fluent      struct {
		Host string
		Port int
		Tag  string
	}

	// ...
}

func init() {
	setEnv()

	cfg = &Config{}
	env.IgnorePrefix()
	err := env.Fill(&cfg)
	if err != nil {
		log.Println("设置env错误：", err)
		panic(err)
	}
}

func setEnv() {
	fi, err := os.Open("./.env")
	if err != nil {
		log.Println("读取.env文件错误：", err)
		panic(err)
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		a, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}

		arr := strings.Split(string(a), "=")
		if 1 == len(arr) {
			continue
		}

		os.Setenv(arr[0], arr[1])
	}
}

func GetEnv() *Config {
	return cfg
}
