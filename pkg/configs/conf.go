package configs

import (
	"errors"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/spf13/viper"
)

var (
	conf   Conf
	_viper *viper.Viper
)

func init() {
	env := os.Getenv("GO_ENV")
	v := viper.New()
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(filename)
	v.AddConfigPath(root)
	v.SetConfigName("config")
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil && env == "" {
		klog.Fatal(err)
	}
	configs := v.AllSettings()
	for key := range configs {
		v.SetDefault(key, configs[key])
	}
	if env != "" {
		file := fmt.Sprintf("config.%s", env)
		v.SetConfigName(file)
		err := v.ReadInConfig()
		if err != nil && !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			klog.Fatal(err)
		} else {
			klog.Infof("successfully read multi environment configuration %s", file)
		}
	}
	err = v.Unmarshal(&conf)
	if err != nil {
		klog.Fatal(err)
	}
	_viper = v
}

func GetConf() *Conf {
	return &conf
}

func GetViper() *viper.Viper {
	return _viper
}

func GetApp() *App {
	return &conf.App
}

func GetDb() *Db {
	return &conf.Db
}

func GetOss() *Oss {
	return &conf.Oss
}

func GetPush() *Push {
	return &conf.Push
}

func GetRegistry() *Registry {
	return &conf.Registry
}
