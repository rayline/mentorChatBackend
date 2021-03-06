package controllers

import (
	"github.com/astaxie/beego"
	"io/ioutil"
	"mentorChatBackend/models/files"
	"mentorChatBackend/models/tokens"
	"mentorChatBackend/models/types"
	"strconv"
)

type FileController struct {
	beego.Controller
}

func (c *FileController) Prepare() {
	TokenString := c.Ctx.GetCookie("token")
	if TokenString == "" {
		return
	} else {
		tokenuint, err := strconv.ParseUint(TokenString, 16, 64)
		if err != nil {
			return
		}
		token := types.TokenID_t(tokenuint)
		uid, err := tokens.Get(token)
		if err != nil {
			return
		}
		c.Data["userid"] = uid
	}
}

func (c *FileController) NewFile() {
	file, _, err := c.GetFile("file")
	if err != nil {
		beego.Error("failed dealing with uploaded file : %v\n", err)
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to recieve uploaded file: " + err.Error(),
		}
		c.ServeJSON()
	} else {
		data, err := ioutil.ReadAll(file)
		if err != nil {
			beego.Error("failed dealing with uploaded file : %v\n", err)
			c.Data["json"] = map[string]interface{}{
				"result": "failed",
				"error":  "failed to read uploaded file content : " + err.Error(),
			}
			c.ServeJSON()
			return
		}
		fileid, err := files.NewFile(data)
		if err != nil {
			beego.Error("failed saving uploaded file : %v\n", err)
			c.Data["json"] = map[string]interface{}{
				"result": "failed",
				"error":  "failed to write uploaded file content : " + err.Error(),
			}
		}
		c.Data["json"] = map[string]interface{}{
			"result": "success",
			"data": map[string]string{
				"fileid": string(fileid),
			},
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
	data, err := files.GetFile(types.FileID_t(Id))
	if err != nil {
		beego.Error("failed dealing with retrieving file : %v\n", err)
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to retrieve file",
		}
	}
	c.Ctx.ResponseWriter.Write(data)
}
