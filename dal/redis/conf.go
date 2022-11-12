package redis

import (
	"context"
	"cybercoin/dal/const_data"
	"cybercoin/util"
	"github.com/go-redis/redis/v8"
	"sync"
)

var Redis *redis.Client
var once = sync.Once{}
var RedisConf Conf

func NewClient(ctx context.Context) *redis.Client {
	once.Do(func() {
		RedisConf = Conf{}
		util.ReadTransfer(const_data.REDIS_CONF, &RedisConf)
		Redis = redis.NewClient(&redis.Options{
			Addr:     RedisConf.Host + ":" + RedisConf.Port,
			Password: RedisConf.Password,
			DB:       RedisConf.DB,
		})
	})
	return Redis
}

type Conf struct {
	Host     string
	Port     string
	Password string
	DB       int
}
