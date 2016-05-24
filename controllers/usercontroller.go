package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"mentorChatBackend/models/tokens"
	"mentorChatBackend/models/types"
	"mentorChatBackend/models/users"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Prepare() {
	TokenString := c.Ctx.GetCookie("token")
	if UIDString == "" {
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

func (c *UserController) AllUsers() {
	if beego.AppConfig.String("runmode") == "dev" {
		if userList, err := users.GetAll(); err != nil {
			c.Data["json"] = map[string]interface{}{
				"result": "failed",
				"error":  "failed to get the list",
			}
			c.ServeJSON()
			return
		} else {
			data, err := json.Marshal(users.GetAll())
			if err != nil {
				beego.BeeLogger.Error("failed to Marshal user list : %v\n", err)
				c.Data["json"] = map[string]interface{}{
					"result": "failed",
					"error":  "failed to generate the list",
				}
				c.ServeJSON()
				return
			}
			c.Ctx.ResponseWriter.Write(data)
		}
	}
}

func (c *UserController) GetUser() {
	requesteeId, err := stringToUserId(c.Ctx.Input.Param("userid"))
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to parse requested ID",
		}
		c.ServeJSON()
		return
	}
	permission := users.NONE
	if uidInterface, exist := c.Data["userid"]; exist == true {
		userid := uidInterface.(types.UserID_t)
		permission = users.GetPermission(userid, requesteeId)
		requestee, err := users.Get(requesteeId)
		if err != nil {
			c.Data["json"] = map[string]interface{}{
				"result": "failed",
				"error":  "failed to get user",
			}
			c.ServeJSON()
			return
		}
		if permission < users.SELF {
			c.Data["json"] = map[string]interface{}{
				"result":      "success",
				"Name":        requestee.Name,
				"Description": requestee.Description,
			}
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = map[string]interface{}{
				"result":      "success",
				"Name":        requestee.Name,
				"Description": requestee.Description,
				"Mail":        requestee.Mail,
			}
			c.ServeJSON()
			return
		}
	}
}

func (c *UserController) ModifyUser() {
	requesteeId, err := stringToUserId(c.Ctx.Input.Param("userid"))
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to parse requested ID",
		}
		c.ServeJSON()
		return
	}
	permission := users.NONE
	if uidInterface, exist := c.Data["userid"]; exist == true {
		userid := uidInterface.(types.UserID_t)
		permission = users.GetPermission(userid, requesteeId)
	}
	requestee, err := users.Get(requesteeId)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to get requested user",
		}
		c.ServeJSON()
		return
	}
	if permission < users.SELF {
		c.Data["json"] = map[string]interface{}{
			"result":      "success",
			"Name":        requestee.Name,
			"Description": requestee.Description,
		}
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = map[string]interface{}{
			"result":      "success",
			"Name":        requestee.Name,
			"Description": requestee.Description,
			"Mail":        requestee.Mail,
		}
		c.ServeJSON()
		return
	}
}

func (c *UserController) LoginUser() {
	passwordStr := c.GetString("password")
	requesteeId, err := stringToUserId(c.Ctx.Input.Param("userid"))
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to parse requested ID",
		}
		c.ServeJSON()
		return
	}
	requestee, err := users.Get(requesteeId)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to get requested user",
		}
		c.ServeJSON()
		return
	}
	if requestee.Validate(passwordStr) {
		c.Ctx.SetCookie("token", tokens.NewToken(requestee.Id))
		c.Data["json"] = map[string]interface{}{
			"result": "success",
		}
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "wrong userid or password",
		}
		c.ServeJSON()
		return
	}
}

func (c *UserController) SendMessage() {
	requesteeId, err := stringToUserId(c.Ctx.Input.Param("userid"))
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to parse requested ID",
		}
		c.ServeJSON()
		return
	}
	permission := users.NONE
	if uidInterface, exist := c.Data["userid"]; exist == true {
		userid := uidInterface.(types.UserID_t)
		permission = users.GetPermission(userid, requesteeId)
	}
	requestee, err := users.Get(requesteeId)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to get requested user",
		}
		c.ServeJSON()
		return
	}
}

func (c *UserController) GetMESSAGE() {

}

func (c *UserController) ModifyFriendList() {

}

func (c *UserController) GetFriendList() {

}

func (c *UserController) SendFriendRequest() {

}

func (c *UserController) GetUserIdByMail() {

}

func (c *UserController) GetUserIdByName() {

}

func stringToUserId(IdStr string) (types.UserID_t, error) {
	IDStr := IdStr
	IdInt, err := strconv.ParseUint(requesteeIDStr, 16, 64)
	ID := types.UserID_t(IdInt)
	return ID, err
}
