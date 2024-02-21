package tool

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	App      App
	Database Database
	Redis    RedisConfig
}

type App struct {
	Name string
	Host string
	Port string
	Mode string
}
type Database struct {
	Driver   string
	User     string
	Password string
	Host     string
	Port     string
	Name     string
	ShowSql  bool
}

type RedisConfig struct {
	Addr     string
	Port     string
	Password string
	Db       int
}

var _cfg Config

func ParseConfig(path string) *Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	err = viper.Unmarshal(&_cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &_cfg
}
