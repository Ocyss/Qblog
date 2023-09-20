package dal

import (
	"github.com/Ocyss/Qblog/cmd/notification/biz/dal/db"
	"github.com/Ocyss/Qblog/cmd/notification/biz/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
