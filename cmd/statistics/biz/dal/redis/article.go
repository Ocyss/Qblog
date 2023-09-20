package redis

import (
	"context"
	"strconv"
	"strings"

	"github.com/redis/go-redis/v9"
)

func GetUvPv(ctx context.Context, aid uint, ua *string) (uv, pv int64, err error) {
	var keyBuf strings.Builder
	keyBuf.WriteString("article:")
	keyBuf.WriteString(strconv.Itoa(int(aid)))
	key := keyBuf.String()

	cmds := rdb.Pipeline()
	cmds.PFAdd(ctx, key, ua)
	cmds.PFCount(ctx, key)
	cmds.Incr(ctx, key)
	results, err := cmds.Exec(ctx)
	if err != nil {
		return
	}
	uv, err = results[1].(*redis.IntCmd).Result()
	if err != nil {
		return
	}
	pv, err = results[2].(*redis.IntCmd).Result()
	return
}
