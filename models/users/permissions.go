package users

import "mentorChatBackend/models/types"

type Permission uint

const (
	ALL         Permission = 10086
	NONE        Permission = 0
	SELF        Permission = 10
	FRIEND      Permission = 5
	STRANGER    Permission = 2
	BLACKLISTED Permission = 1 //TODO: implement blacklist
)

func GetPermission(RequesterID types.UserID_t, RequesteeID types.UserID_t) Permission {
	if RequesterID == 0 {
		return ALL
	}
	if RequesterID == RequesteeID {
		return SELF
	}
	if requester, err := GetWithNoInformation(RequesterID); err != nil {
		return NONE
	} else if isfriend := requester.IsFriend(RequesteeID); err == nil || isfriend == true {
		return FRIEND
	}
	if _, err := GetWithNoInformation(RequesteeID); err != nil {
		return NONE
	}
	return STRANGER
}
