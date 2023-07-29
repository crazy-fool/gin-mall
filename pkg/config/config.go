package config

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var vp *viper.Viper

func init() {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		flag.StringVar(&envConf, "conf", "config/local.yml", "config path, eg: -conf config/local.yml")
		flag.Parse()
	}
	if envConf == "" {
		envConf = "local"
	}
	fmt.Println("load conf file:", envConf)
	vp = getConfig(envConf)
}

func GetConfig() *viper.Viper {
	return vp
}

func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
