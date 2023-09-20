package dal

import (
	"github.com/Ocyss/Qblog/cmd/user/biz/dal/db"
	"github.com/Ocyss/Qblog/cmd/user/biz/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
