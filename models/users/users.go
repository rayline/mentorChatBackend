package users

import "mentorChatBackend/models/types"

type User struct {
	Id       types.UserID_t
	Password types.Password_t
	Name     string
}

func GetAll() []user.User {

}

func Get(Id types.UserID_t) user.User {

}

func GetByName(Name string) user.User {

}

func GetByMail(Mail string) user.User {

}

func Set(Id types.UserID_t, u user.User) {

}
