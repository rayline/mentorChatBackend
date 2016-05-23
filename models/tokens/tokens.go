package tokens

import "time"
import "mentorChatBackend/models/types"
import "sync"
import "fmt"

var tokenMap = map[types.TokenID_t]Token{}
var tokenCnt types.TokenID_t = 1

var tokenMutex sync.RWMutex

const MESSAGE_BLOCKING_LIMIT time.Duration = 30 * time.Second

func init() {

}

type token struct {
	Id       uint64
	lastUsed time.Time
}

func NewToken(Id types.UserID_t) types.TokenID_t {
	tokenMutex.Lock()
	defer tokenMutex.Unlock()
	t := new(token)
	t.Id = Id
	t.lastUsed = time.Now()
	tokenCnt++
	tid := tokenCnt
	tokenMap[tid] = t
	return tid
}

func Get(Id types.TokenID_t) (types.UserID_t, error) {
	tokenMutex.RLock()
	defer tokenMutex.RUnlock()
	var t token
	if t, exist = tokenMap[Id]; exist == false {
		return 0, fmt.Errorf("Token: Token ID %d does not exist", Id)
	}
	if t.lastUsed.Add(MESSAGE_BLOCKING_LIMIT).Before(time.Now()) {
		return 0, fmt.Errorf("Token: Token ID %d is expired", Id)
	}
	return t.Id, nil
}
