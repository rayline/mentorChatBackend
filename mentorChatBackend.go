package main

import (
	"github.com/astaxie/beego"
	_ "mentorChatBackend/routers"
)

func main() {
	beego.SetStaticPath("/", "static")
	beego.BConfig.Listen.EnableHTTPS = true
	beego.BConfig.Listen.HTTPSCertFile = beego.AppConfig.String("HTTPSCertFile")
	beego.BConfig.Listen.HTTPSKeyFile = beego.AppConfig.String("HTTPSKeyFile")
	beego.Run()
}
