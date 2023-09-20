package redis

import (
	"context"
	"time"

	"github.com/Ocyss/Qblog/pkg/configs"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/redis/go-redis/v9"
)

var rdb *redis.Client

func Init() {
	conf := configs.GetDb().Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     conf.Address,
		Password: conf.Password,
		DB:       conf.Db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		klog.Fatal(err)
	}
}
