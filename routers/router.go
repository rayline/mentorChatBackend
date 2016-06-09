package routers

import (
	"github.com/astaxie/beego"
	"mentorChatBackend/controllers"
)

func init() {
	beego.Router("/api/users", &controllers.UserController{}, "get:AllUsers")
	beego.Router("/api/user/new", &controllers.UserController{}, "get:NewUser")
	beego.Router("/api/user/logout", &controllers.UserController{}, "get:LogoutUser")
	beego.Router("/api/user/:userid", &controllers.UserController{}, "get:GetUser;post:ModifyUser")
	beego.Router("/api/user/:userid/login", &controllers.UserController{}, "post:LoginUser")
	beego.Router("/api/user/:userid/message", &controllers.UserController{}, "post:SendMessage;get:GetMESSAGE")
	beego.Router("/api/user/:userid/friendlist", &controllers.UserController{}, "post:ModifyFriendList;get:GetFriendList")
	beego.Router("/api/user/:userid/friendrequest", &controllers.UserController{}, "post:SendFriendRequest")
	beego.Router("/api/usermail/:usermail", &controllers.UserController{}, "get:GetUserIdByMail")
	beego.Router("/api/username/:username", &controllers.UserController{}, "get:GetUserIdByName")
	beego.Router("/api/file/new", &controllers.FileController{}, "post:NewFile")
	beego.Router("/api/file/:fileid", &controllers.FileController{}, "get:RetrieveFile")
	beego.Router("/", &controllers.MainController{})
}
