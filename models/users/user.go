package users

import "github.com/astaxie/beego"
import "mentorChatBackend/models/types"
import "github.com/garyburd/redigo/redis"
import "mentorChatBackend/models/consts"
import "encoding/json"
import "time"

type User struct {
	Id                      types.UserID_t
	Password                types.Password_t
	Name, Description, Mail string
}

func newPool(server, password string, database int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			c.Do("SELECT", database)
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

var (
	pool0, pool1, pool2 *redis.Pool
	redisServer         = consts.RedisServer
	redisPassword       = consts.RdisPassword
)

func init() {
	pool0 = newPool(redisServer, redisPassword, 0)
	pool1 = newPool(redisServer, redisPassword, 1)
	pool2 = newPool(redisServer, redisPassword, 2)
}

func (u *User) IsFriend(Id types.UserID_t) bool {
	conn := pool2.Get()
	defer conn.Close()
	is, err := redis.Bool(conn.Do("SISMEMBER", u.Id, Id))
	if err != nil {
		beego.BeeLogger.Error("%v\n", err)
	}
	return is
}

func (u *User) AcceptAsFriend(Id types.UserID_t) {
	conn := pool2.Get()
	defer conn.Close()
	_, err := conn.Do("SADD", u.Id, Id)
	if err != nil {
		beego.BeeLogger.Error("%v\n", err)
	}
	_, err = conn.Do("SADD", Id, u.Id)
	if err != nil {
		beego.BeeLogger.Error("%v\n", err)
	}
}

func (u *User) DeleteFriend(Id types.UserID_t) {
	conn := pool2.Get()
	defer conn.Close()
	_, err := conn.Do("SREM", u.Id, Id)
	if err != nil {
		beego.BeeLogger.Error("%v\n", err)
	}
	_, err = conn.Do("SREM", Id, u.Id)
	if err != nil {
		beego.BeeLogger.Error("%v\n", err)
	}
}

func (u *User) AddMESSAGE(MESSAGE types.Message_t) {
	conn := pool1.Get()
	defer conn.Close()
	data, err := json.Marshal(MESSAGE)
	if err != nil {
		beego.BeeLogger.Error("%v\n", err)
		return
	}
	if _, err = conn.Do("LPUSH", data); err != nil {
		beego.BeeLogger.Error("%v\n", err)
	}
}
