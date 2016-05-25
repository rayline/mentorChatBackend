package users

import "mentorChatBackend/models/types"
import "github.com/garyburd/redigo/redis"
import "fmt"
import "encoding/json"
import "strconv"
import "log"

func GetAll() ([]*User, error) {
	conn := pool0.Get()
	defer conn.Close()
	Idstrings, err := redis.Strings(conn.Do("KEYS", "[0-9]*"))
	if err != nil {
		log.Printf("%v\n", err)
		return nil, fmt.Errorf("users: failed getting user list:%v\n", err)
	}
	uList := []*User{}
	for _, v := range Idstrings {
		Id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			log.Printf("Unexpected number conversion error: %v", err)
			continue
		}
		u, err := Get(types.UserID_t(Id))
		if err != nil {
			log.Printf("%v\n", err)
			continue
		}
		uList = append(uList, u)
	}
	return uList, nil
}

func GetAllUserId() ([]types.UserID_t, error) {
	conn := pool0.Get()
	defer conn.Close()
	Idstrings, err := redis.Strings(conn.Do("KEYS", "[0-9]*"))
	if err != nil {
		log.Printf("%v\n", err)
		return nil, fmt.Errorf("users: failed getting user list:%v\n", err)
	}
	uList := []types.UserID_t{}
	for _, v := range Idstrings {
		Id, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			continue
		}
		uList = append(uList, types.UserID_t(Id))
	}
	return uList, nil
}

func Get(Id types.UserID_t) (*User, error) {
	conn := pool0.Get()
	defer conn.Close()
	originJson, err := redis.Bytes(conn.Do("GET", Id))
	if err == redis.ErrNil {
		return nil, fmt.Errorf("users: user %v does not exist", Id)
	} else if err != nil {
		log.Printf("%v\n", err)
		return nil, fmt.Errorf("users: failed getting user with uid %v\n%v\n", Id, err)
	}
	var u User
	json.Unmarshal(originJson, &u)
	return &u, nil
}

func GetByName(Name string) (*User, error) {
	conn := pool4.Get()
	defer conn.Close()
	Id, err := redis.Uint64(conn.Do("GET", Name))
	if err == redis.ErrNil {
		return nil, fmt.Errorf("users: user name %v does not exist", Name)
	} else if err != nil {
		log.Printf("%v\n", err)
		return nil, fmt.Errorf("users: failed getting user with name %v\n%v\n", Name, err)
	}
	return Get(types.UserID_t(Id))
}

func GetByMail(Mail string) (*User, error) {
	conn := pool3.Get()
	defer conn.Close()
	Id, err := redis.Uint64(conn.Do("GET", Mail))
	if err == redis.ErrNil {
		return nil, fmt.Errorf("users: user mail %v does not exist", Mail)
	} else if err != nil {
		log.Printf("%v\n", err)
		return nil, fmt.Errorf("users: failed getting user with mail %v\n%v\n", Mail, err)
	}
	return Get(types.UserID_t(Id))
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
	if err != nil && err != redis.ErrNil {
		log.Printf("%v\n", err)
	}
	var origin User
	json.Unmarshal(originJson, &origin)
	if origin.Id == 0 {
		origin.Id = Id
	}
	if origin.Id != Id {
		log.Printf("ID stored in DB not same with input\n")
	}
	if u.Password != "" {
		origin.Password = u.Password
	}
	if u.Name != "" {
		if origin.Name != u.Name {
			setName(Id, u.Name)
		}
		origin.Name = u.Name
	}
	if u.Description != "" {
		origin.Description = origin.Description
	}
	if u.Mail != "" {
		if origin.Mail != u.Mail {
			setMail(Id, u.Mail)
		}
		origin.Mail = u.Mail
	}
	data, _ := json.Marshal(origin)
	if _, err := conn.Do("SET", Id, data); err != nil {
		log.Printf("%v\n", err)
		return
	}
}

func AllocUID() types.UserID_t {
	UIDAllocMutex.Lock()
	defer UIDAllocMutex.Unlock()
	defer func() {
		nextUID++
		conn := pool0.Get()
		defer conn.Close()
		conn.Do("SET", "NEXTUID", nextUID)
	}()
	u := User{}
	u.Id = nextUID
	Set(nextUID, u)
	return nextUID
}

func setName(Id types.UserID_t, name string) {
	conn := pool4.Get()
	defer conn.Close()
	if _, err := conn.Do("SET", Id, name); err != nil {
		log.Printf("%v\n", err)
		return
	}
}

func setMail(Id types.UserID_t, mail string) {
	conn := pool3.Get()
	defer conn.Close()
	if _, err := conn.Do("SET", Id, mail); err != nil {
		log.Printf("%v\n", err)
		return
	}
}
