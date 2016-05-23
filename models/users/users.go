package users

import "github.com/astaxie/beego"
import "mentorChatBackend/models/types"
import "github.com/garyburd/redigo/redis"
import "fmt"
import "encoding/json"

func GetAll() []User {
	return nil
}

func Get(Id types.UserID_t) (*User, error) {
	return nil, nil
}

func GetByName(Name string) (*User, error) {
	return nil, nil
}

func GetByMail(Mail string) (*User, error) {
	return nil, nil
}

func GetWithNoInformation(Id types.UserID_t) (*User, error) {
	u := &User{}
	u.Id = Id
	conn := pool0.Get()
	defer conn.Close()
	if exist, err := redis.Bool(conn.Do("EXISTS", Id)); err != nil || exist == false {
		return nil, fmt.Errorf("No such user %d", Id)
	}
	return u, nil
}

func Set(Id types.UserID_t, u User) {
	conn := pool0.Get()
	defer conn.Close()
	originJson, err := redis.Bytes(conn.Do("GET", Id))
	if err != nil {
		beego.BeeLogger.Error("%v\n", err)
		return
	}
	var origin User
	json.Unmarshal(originJson, &origin)
	if u.Password != nil && u.Password != "" {
		origin.Password = u.Password
	}
	if u.Name != nil && u.Name != "" {
		if origin.Name != nil && origin.Name != u.Name {
			setName(Id, u.Name)
		}
		origin.Name = u.Name
	}
	if u.Description != nil && u.Description != "" {
		origin.Description = origin.Description
	}
	if u.Mail != nil && u.Mail != "" {
		if origin.Mail != nil && origin.Mail != u.Mail {
			setName(Id, u.Mail)
		}
		origin.Mail = origin.Mail
	}
}

func setName(Id types.UserID_t, name string) {

}

func setMail(Id types.UserID_t, mail string) {

}
