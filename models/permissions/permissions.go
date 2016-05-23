package permissions

import "mentorChatBackend/models/types"
import "mentorChatBackend/models/users"

type Permission uint

const (
	ALL      Permission = 2147483647
	NONE     Permission = 0
	SELF     Permission = 10
	FRIEND   Permission = 5
	STRANGER Permission = 1
)

func Get(RequesterID types.UserID_t, RequesteeID types.UserID_t) Permission {
	if RequesterID == 0 {
		return ALL
	}
	if RequesterID == RequesteeID {
		return SELF
	}
	if requester, err := users.GetWithNoInformation(RequesterID); err != nil {
		return NONE
	}
	if isfriend, err = requester.IsFriend(RequesteeID); err == nil || isfriend == true {
		return FRIEND
	}
	if requestee, err := users.GetWithNoInformation(RequesteeID); err != nil {
		return NONE
	}
	return STRANGER
}
