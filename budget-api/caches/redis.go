package caches

import (
	"budget-api/config"

	"github.com/go-redis/redis"
	toml "github.com/pelletier/go-toml"
)

var (
	// Cache 是整个包返回的内容，即一个redis client
	Cache = New()
)

// New 返回一个redis client的单例实例
func New() *redis.Client {
	configTree := config.Conf.Get("redis").(*toml.Tree)
	return redis.NewClient(&redis.Options{
		Addr:     configTree.Get("Addr").(string),
		Password: configTree.Get("Password").(string), // no password set
		DB:       int(configTree.Get("DB").(int64)),   // 因为系统是64位的，所以默认的 int 型是 int64
	})
}
