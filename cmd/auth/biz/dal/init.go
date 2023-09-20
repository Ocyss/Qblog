package dal

import (
	"github.com/Ocyss/Qblog/cmd/auth/biz/dal/db"
	"github.com/Ocyss/Qblog/cmd/auth/biz/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
