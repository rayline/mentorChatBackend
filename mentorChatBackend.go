package main

import (
	"github.com/astaxie/beego"
	"log"
	_ "mentorChatBackend/routers"
	"os"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := os.OpenFile("mentorChat.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Unable to create log file : ", err)
	}
	log.SetOutput(f)
}

func main() {
	beego.Run()
}
