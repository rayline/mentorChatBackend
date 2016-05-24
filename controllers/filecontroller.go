package controllers

import (
	"github.com/astaxie/beego"
	"mentorChatBackend/models/files"
	"mentorChatBackend/models/types"
	"mentorChatBackend/models/users"
	"mentorChatBackend/models/tokens"
	"strconv"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) Prepare() {
	TokenString := c.Ctx.GetCookie("token")
	if UIDString==""{
		return
	}else{
		tokenuint,err := strconv.ParseUint(TokenString, 16, 64)
		if err!=nil{
			return
		}
		token := types.TokenID_t(tokenuint)
		uid,err := tokens.Get(token)
		if err!=nil{
			return
		}
		c.Data["userid"] = uid
	}
}

func (c *FileController) NewFile() {
	file, fileheader, err := c.GetFile("file")
	if err != nil {
		beego.BeeLogger.Error("failed dealing with uploaded file : %v\n", err)
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to upload file",
		}
		c.ServeJSON()
	} else {
		c.Data["json"] = map[string]interface{}{
			"result": "success",
		}
		c.ServeJSON()
	}
}

func (c *FileController) RetrieveFile() {
	Id := c.Ctx.Input.Param(":fileid")
	if Id == "" {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "No fileid input",
		}
	}
	data, err := files.GetFile(Id)
	if err != nil {
		beego.BeeLogger.Error("failed dealing with retrieving file : %v\n", err)
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to retrieve file",
		}
	}
	c.
}
