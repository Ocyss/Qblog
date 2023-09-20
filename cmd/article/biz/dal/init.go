package dal

import (
	"github.com/Ocyss/Qblog/cmd/article/biz/dal/db"
	"github.com/Ocyss/Qblog/cmd/article/biz/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
