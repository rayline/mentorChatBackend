package controllers

import (
	"github.com/astaxie/beego"
	"mentorChatBackend/models/files"
	"mentorChatBackend/models/types"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) NewFile() {

}

func (c *FileController) GetFile() {
	Id := c.GetString("fileid")
	if Id==nil||Id==""{
		c.
	}
	files.GetFile(Id)
}

func (c *FileController) DeleteFile() {

}
