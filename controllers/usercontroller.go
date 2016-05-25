package controllers

import (
	"fmt"
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
	if TokenString == "" {
		return
	} else {
		tokenuint, err := strconv.ParseUint(TokenString, 10, 64)
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
		if data, err := users.GetAllUserId(); err != nil {
			c.Data["json"] = map[string]interface{}{
				"result": "failed",
				"error":  "failed to get the list",
			}
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = map[string]interface{}{
				"result": "success",
				"data":   data,
			}
			c.ServeJSON()
			return
		}
	}
}

func (c *UserController) GetUser() {
	requesteeId, err := stringToUserId(c.Ctx.Input.Param(":userid"))
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
				"result": "success",
				"data": map[string]string{
					"Name":        requestee.Name,
					"Description": requestee.Description,
				},
			}
			c.ServeJSON()
			return
		} else {
			c.Data["json"] = map[string]interface{}{
				"result": "success",
				"data": map[string]string{
					"Name":        requestee.Name,
					"Description": requestee.Description,
					"Mail":        requestee.Mail,
				},
			}
			c.ServeJSON()
			return
		}
	} else {
		requestee, err := users.Get(requesteeId)
		if err != nil {
			c.Data["json"] = map[string]interface{}{
				"result": "failed",
				"error":  "failed to get user",
			}
			c.ServeJSON()
			return
		}
		c.Data["json"] = map[string]interface{}{
			"result": "success",
			"data": map[string]string{
				"Name":        requestee.Name,
				"Description": requestee.Description,
			},
		}
		c.ServeJSON()
		return
	}
}

func (c *UserController) ModifyUser() {
	//TODO:it is not checked whether a name a duplicated and when changing name the index of old name was not removed
	requesteeId, err := stringToUserId(c.Ctx.Input.Param(":userid"))
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to parse requested ID",
		}
		c.ServeJSON()
		return
	}
	permission := users.NONE
	userid := types.UserID_t(0)
	if uidInterface, exist := c.Data["userid"]; exist == true {
		userid = uidInterface.(types.UserID_t)
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
			"result": "failed",
			"error":  fmt.Sprintf("Access Denied %v to %v", userid, requesteeId),
		}
		c.ServeJSON()
		return
	} else {
		if name := c.GetString("name"); name != "" {
			requestee.Name = name
		}
		if mail := c.GetString("mail"); mail != "" {
			requestee.Mail = mail
		}
		if description := c.GetString("description"); description != "" {
			requestee.Description = description
		}
		if password := c.GetString("password"); password != "" {
			requestee.Password = types.Password_t(password)
		}
		users.Set(requestee.Id, *requestee)
		c.Data["json"] = map[string]interface{}{
			"result": "success",
		}
		c.ServeJSON()
		return
	}
}

func (c *UserController) LoginUser() {
	passwordStr := c.GetString("password")
	requesteeId, err := stringToUserId(c.Ctx.Input.Param(":userid"))
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
	if requestee.Validate(types.Password_t(passwordStr)) {
		c.Ctx.SetCookie("token", string(tokens.NewToken(requestee.Id)))
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

func (c *UserController) NewUser() {
	Id := users.AllocUID()
	c.Data["json"] = map[string]interface{}{
		"result": "success",
		"data": map[string]interface{}{
			"userid": Id,
		},
	}
	token := tokens.NewToken(Id)
	c.Ctx.SetCookie("token", strconv.FormatUint(uint64(token), 10))
	c.ServeJSON()
}

func (c *UserController) SendMessage() {
	requesteeId, err := stringToUserId(c.Ctx.Input.Param(":userid"))
	message := c.GetString("message")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to parse requested ID",
		}
		c.ServeJSON()
		return
	}
	userid := types.UserID_t(0)
	if uidInterface, exist := c.Data["userid"]; exist == true {
		userid = uidInterface.(types.UserID_t)
	}
	requester, err := users.GetWithNoInformation(userid)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to get requester",
		}
		c.ServeJSON()
		return
	}
	err = requester.SendMessage(requesteeId, message)
	if err == users.ErrAccessDenied {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "access denied",
		}
		c.ServeJSON()
		return
	} else if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "Server Internal Error",
		}
		c.ServeJSON()
		return
	}
}

func (c *UserController) GetMESSAGE() {
	requesteeId, err := stringToUserId(c.Ctx.Input.Param(":userid"))
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
			"result": "failed",
			"error":  "access denied",
		}
		c.ServeJSON()
		return
	}
	MESSAGE := requestee.GetMESSAGE()
	if MESSAGE == nil {
		c.Data["json"] = map[string]interface{}{
			"result": "success",
		}
		c.ServeJSON()
		return
	} else {
		c.Data["json"] = map[string]interface{}{
			"result": "success",
			"data":   *MESSAGE,
		}
		c.ServeJSON()
		return
	}
}

func (c *UserController) ModifyFriendList() {
	if c.GetString("behavior") != "DELETE" {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "Unknown Behavior",
		}
		c.ServeJSON()
		return
	}
	user := &users.User{}
	if uidInterface, exist := c.Data["userid"]; exist == true {
		userid := uidInterface.(types.UserID_t)
		var err error
		user, err = users.GetWithNoInformation(userid)
		if err != nil {
			c.Data["json"] = map[string]interface{}{
				"result": "failed",
				"error":  "Access Denied",
			}
		}
	}
	UIDStrings := c.GetStrings("friendlist")
	for _, v := range UIDStrings {
		if userid, err := stringToUserId(v); err != nil {
			continue
		} else {
			user.DeleteFriend(userid)
		}
	}
}

func (c *UserController) GetFriendList() {
	user := &users.User{}
	if uidInterface, exist := c.Data["userid"]; exist == true {
		userid := uidInterface.(types.UserID_t)
		var err error
		user, err = users.GetWithNoInformation(userid)
		if err != nil {
			c.Data["json"] = map[string]interface{}{
				"result": "failed",
				"error":  "Access Denied",
			}
			c.ServeJSON()
			return
		}
	}
	uList := user.GetFriendList()
	c.Data["json"] = map[string]interface{}{
		"result": "success",
		"data": map[string]interface{}{
			"friendlist": uList,
		},
	}
	c.ServeJSON()
	return
}

func (c *UserController) SendFriendRequest() {
	requesteeId, err := stringToUserId(c.Ctx.Input.Param(":userid"))
	message := c.GetString("message")
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to parse requested ID",
		}
		c.ServeJSON()
		return
	}
	userid := types.UserID_t(0)
	if uidInterface, exist := c.Data["userid"]; exist == true {
		userid = uidInterface.(types.UserID_t)
	}
	requester, err := users.GetWithNoInformation(userid)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "failed to get requester",
		}
		c.ServeJSON()
		return
	}
	err = requester.SendFriendRequest(requesteeId, message)
	if err == users.ErrAccessDenied {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "access denied",
		}
		c.ServeJSON()
		return
	} else if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "Server Internal Error",
		}
		c.ServeJSON()
		return
	}
}

func (c *UserController) GetUserIdByMail() {
	mail := c.Ctx.Input.Param(":usermail")
	u, err := users.GetByMail(mail)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "No such mail : " + mail + err.Error(),
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{
		"result": "success",
		"data": map[string]interface{}{
			"userid": u.Id,
		},
		"dbgdata": u,
	}
	c.ServeJSON()
	return
}

func (c *UserController) GetUserIdByName() {
	name := c.Ctx.Input.Param(":username")
	u, err := users.GetByName(name)
	if err != nil {
		c.Data["json"] = map[string]interface{}{
			"result": "failed",
			"error":  "No such name :" + name + err.Error(),
		}
		c.ServeJSON()
		return
	}
	c.Data["json"] = map[string]interface{}{
		"result": "success",
		"data": map[string]interface{}{
			"userid": u.Id,
		},
	}
	c.ServeJSON()
	return
}

func stringToUserId(IdStr string) (types.UserID_t, error) {
	requesteeIDStr := IdStr
	IdInt, err := strconv.ParseUint(requesteeIDStr, 10, 64)
	ID := types.UserID_t(IdInt)
	return ID, err
}
