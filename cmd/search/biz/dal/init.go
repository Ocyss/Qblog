package dal

import (
	"github.com/Ocyss/Qblog/cmd/search/biz/dal/db"
	"github.com/Ocyss/Qblog/cmd/search/biz/dal/redis"
)

func Init() {
	db.Init()
	redis.Init()
}
