package users

import "mentorChatBackend/models/types"
import "github.com/garyburd/redigo/redis"
import "mentorChatBackend/models/consts"
import "encoding/json"
import "time"
import "sync"
import "fmt"
import "strconv"
import "log"
import "os"

type User struct {
	Id                      types.UserID_t
	Password                types.Password_t
	Name, Description, Mail string
}

func newPool(server, password string, database int) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     100,
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
	pool0, pool1, pool2, pool3, pool4 *redis.Pool
	redisServer                                      = consts.RedisServer
	redisPassword                                    = consts.RdisPassword
	nextUID                           types.UserID_t = 1
	UIDAllocMutex                     sync.Mutex
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	f, err := os.OpenFile("mentorChatUSERS.log", os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("Unable to create log file : ", err)
	}
	log.SetOutput(f)
	pool0 = newPool(redisServer, redisPassword, 0)
	pool1 = newPool(redisServer, redisPassword, 1)
	pool2 = newPool(redisServer, redisPassword, 2)
	pool3 = newPool(redisServer, redisPassword, 3)
	pool4 = newPool(redisServer, redisPassword, 4)
	conn := pool0.Get()
	defer conn.Close()
	id, err := redis.Uint64(conn.Do("GET", "NEXTUID"))
	if err == redis.ErrNil {
		conn.Do("SET", "NEXTUID", 1)
	} else if err != nil {
		panic(fmt.Sprintf("Failed to get uid count: %v", err))
	}
	nextUID = types.UserID_t(id)
}

func (u *User) IsFriend(Id types.UserID_t) bool {
	conn := pool2.Get()
	defer conn.Close()
	is, err := redis.Bool(conn.Do("SISMEMBER", uint64(u.Id), uint64(Id)))
	if err != nil {
		log.Printf("%v\n", err)
	}
	return is
}

func (u *User) AcceptAsFriend(Id types.UserID_t) {
	conn := pool2.Get()
	defer conn.Close()
	_, err := conn.Do("SADD", uint64(u.Id), uint64(Id))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

func (u *User) DeleteFriend(Id types.UserID_t) {
	conn := pool2.Get()
	defer conn.Close()
	_, err := conn.Do("SREM", uint64(u.Id), uint64(Id))
	if err != nil {
		log.Printf("%v\n", err)
	}
	_, err = conn.Do("SREM", uint64(Id), uint64(u.Id))
	if err != nil {
		log.Printf("%v\n", err)
	}
}

func (u *User) AddMESSAGE(MESSAGE types.Message_t) {
	conn := pool1.Get()
	defer conn.Close()
	data, err := json.Marshal(MESSAGE)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	if _, err = conn.Do("LPUSH", uint64(u.Id), data); err != nil {
		log.Printf("%v\n", err)
	}
}

var ErrAccessDenied = fmt.Errorf("Acess Denied")

const (
	UserMessage       = "U"
	SystemAnnouncment = "S"
	FriendRequest     = "F"
)

func (u *User) SendMessage(Id types.UserID_t, message string) error {
	if GetPermission(u.Id, Id) >= FRIEND {
		requestee, err := GetWithNoInformation(Id)
		if err != nil {
			return nil
		}
		requestee.AddMESSAGE(types.Message_t{
			Source:  u.Id,
			Content: message,
			Type:    UserMessage,
		})
		return nil
	} else {
		return ErrAccessDenied
	}
}

func (u *User) SendFriendRequest(Id types.UserID_t, message string) error {
	if GetPermission(u.Id, Id) >= STRANGER {
		requestee, err := GetWithNoInformation(Id)
		if err != nil {
			return nil
		}
		requestee.AddMESSAGE(types.Message_t{
			Source:  u.Id,
			Content: message,
			Type:    FriendRequest,
		})
		//because it is really complicated, we now take the friend system that one will acknowledge the requestee as friend as soon as he sends the request. One can clear the relationship easily by deleting friend though not really friends.
		u.AcceptAsFriend(Id)
		return nil
	} else {
		return ErrAccessDenied
	}
}

func (u *User) SendSystemAnnouncment(message string) error {
	u.AddMESSAGE(types.Message_t{
		Source:  0,
		Content: message,
		Type:    SystemAnnouncment,
	})
	return nil
}

func (u *User) Validate(password types.Password_t) bool {
	if u.Password == password {
		return true
	} else {
		return false
	}
}

const BLOCK_TIME = "15"

func (u *User) GetMESSAGE() *types.Message_t {
	conn := pool1.Get()
	defer conn.Close()
	data, err := redis.MultiBulk(conn.Do("BRPOP", uint64(u.Id), BLOCK_TIME))
	if err == redis.ErrNil {
		return nil
	} else if err == nil {
		MESSAGEBytes, err := redis.Bytes(data[1], nil)
		if err != nil {
			log.Printf("Failed to parse message returned : %v\n", err)
			return nil
		}
		MESSAGE := types.Message_t{}
		err = json.Unmarshal(MESSAGEBytes, &MESSAGE)
		if err != nil {
			log.Printf("Failed to parse message returned : %v\n", err)
			return nil
		}
		return &MESSAGE
	} else {
		log.Printf("Failed to retrieve message : %v\n", err)
		return nil
	}
}

func (u *User) GetFriendList() []types.UserID_t {
	conn := pool2.Get()
	defer conn.Close()
	Idstrings, err := redis.Strings(conn.Do("SMEMBERS", uint64(u.Id)))
	if err != nil {
		log.Printf("%v\n", err)
		return nil
	}
	uList := []types.UserID_t{}
	for _, v := range Idstrings {
		Id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			continue
		}
		uList = append(uList, types.UserID_t(Id))
	}
	return uList
}
