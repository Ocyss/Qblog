package configs

import (
	"testing"
)

func TestGetConf(t *testing.T) {
	conf := GetViper()
	t.Log(conf.AllSettings())
	t.Log("Pgsql:", GetDb().Pgsql)
	t.Log("Oss:", GetOss())
}
