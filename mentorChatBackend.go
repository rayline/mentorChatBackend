package main

import (
	"github.com/astaxie/beego"
	"log"
	_ "mentorChatBackend/routers"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func main() {
	beego.Run()
}
