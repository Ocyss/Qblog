package dal

import (
	"github.com/Ocyss/Qblog/cmd/storage/biz/dal/db"
	"github.com/Ocyss/Qblog/cmd/storage/biz/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
