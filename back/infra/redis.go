package infra

import "github.com/go-redis/redis/v8"

var RS *redis.Client

func init() {
	RS = redis.NewClient(&redis.Options{
		Addr:     "", // Redis服务器地址和端口
		Password: "",               // Redis密码，如果没有设置则为空字符串
		DB:       0,                // 使用默认数据库
	})
}

