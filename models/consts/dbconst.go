package consts

import "github.com/astaxie/beego"

var (
	RedisServer   = "127.0.0.1:6379"
	RedisPassword = "redisPassword"
)

func init() {
	RedisServer = beego.AppConfig.String("RedisServer")
	RedisPassword = beego.AppConfig.String("RedisPassword")
}
