package main

import (
	"github.com/astaxie/beego"
	"log"
	_ "mentorChatBackend/routers"
	"os"
)

func init() {
}

func main() {
	beego.Run()
}
