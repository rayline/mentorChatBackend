package tokens

import "time"
import "mentorChatBackend/models/types"

type Token struct {
	Id       uint64
	lastUsed time.Time
}

func (t *Token) New(Id types.UserID_t) {
	t = new(Token)
	t.Id = Id
	t.lastUsed = time.Now()
}
