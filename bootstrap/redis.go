package bootstrap

import (
	"gorbac/pkg/redis"
)

// SetupRedis 初始化Redis
func SetupRedis(selectDB int) {
	redis.Client = new(redis.Redis).ConnectDB(selectDB)
}

// RedisClose 关闭redis
func RedisClose() {
	redis.Client.Close()
}
