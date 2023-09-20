package dal

import (
	"github.com/Ocyss/Qblog/cmd/statistics/biz/dal/db"
	"github.com/Ocyss/Qblog/cmd/statistics/biz/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
