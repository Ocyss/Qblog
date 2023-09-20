package dal

import (
	"github.com/Ocyss/Qblog/cmd/comment/biz/dal/db"
	"github.com/Ocyss/Qblog/cmd/comment/biz/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
