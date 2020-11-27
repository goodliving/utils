package utils

import "github.com/shima-park/agollo"

type MysqlConfig struct {
	Host string
	DbName string
	User string
	Password string
}

const (
	FlagMysqlUser = "mysql.user"
	FlagMysqlPassword = "mysql.password"
	FlagMysqlHost = "mysql.host"
	FlagMysqlDbName= "mysql.db"
)

func GetMysqlInfo() MysqlConfig {
	return MysqlConfig{
		User:     agollo.Get(FlagMysqlUser),
		Password: agollo.Get(FlagMysqlPassword),
		Host:     agollo.Get(FlagMysqlHost),
		DbName:   agollo.Get(FlagMysqlDbName),
	}
}