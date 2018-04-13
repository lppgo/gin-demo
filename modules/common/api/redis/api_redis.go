package redis

import (
	"demo/services/hsredis"
)

var (
	defaultCache = "coin"   //TODO::配置优化
)

func GetRedisCache() *hsredis.Redis {
	return getHandle(defaultCache)
}

func getHandle(name string) *hsredis.Redis {
	var redisCmdClient *hsredis.Redis
	redisCmdClient = hsredis.GetInstance()
	redisCmdClient.SetCon(name)
	return redisCmdClient
}
