package tests

import "testing"
import "mentorChatBackend/models/users"
import "mentorChatBackend/models/types"

func TestUserAlloc(t *testing.T) {
	//testing basic user registering
	uid := users.AllocUID()
	if uid == 0 {
		t.Fatalf("Failed to alloc user")
	}
	u, err := users.Get(uid)
	if err != nil {
		t.Fatal(err)
	}
	u.Password = "123456"
	users.Set(uid, u)
	u, err = users.Get(uid)
	if err != nil {
		t.Fatal(err)
	}
	if u.Password != "123456" {
		t.Fatal("Failed to change password")
	}

}

func TestUserChangeName(t *testing.T) {
	//tesing user name change and indexing functions
	uid := users.AllocUID()
	if uid == 0 {
		t.Fatalf("Failed to alloc user")
	}
	u, err := users.Get(uid)
	if err != nil {
		t.Fatal(err)
	}
	u.Name = "wyr"
	u.Password = "123456"
	users.Set(uid, u)
	u, err = users.Get(uid)
	if err != nil {
		t.Fatal(err)
	}
	if u.Name != "wyr" {
		t.Fatal("Failed to change password")
	}
	uidb, err := users.GetByName("wyr")
	if err != nil {
		t.Fatal(err)
	}
	if uidb != uid {
		t.Fatal("failed to index")
	}
}

func TestUserChangeMail(t *testing.T) {
	//tesing user name change and indexing functions
	uid := users.AllocUID()
	if uid == 0 {
		t.Fatalf("Failed to alloc user")
	}
	u, err := users.Get(uid)
	if err != nil {
		t.Fatal(err)
	}
	u.Mail = "wyr"
	u.Password = "123456"
	users.Set(uid, u)
	u, err = users.Get(uid)
	if err != nil {
		t.Fatal(err)
	}
	if u.Mail != "wyr" {
		t.Fatal("Failed to change password")
	}
	uidb, err := users.GetByMail("wyr")
	if err != nil {
		t.Fatal(err)
	}
	if uidb != uid {
		t.Fatal("failed to index")
	}
}

func TestUserGetMESSAGE(t *testing.T) {
	//tesing system adding message to message queue and user retrieving it
	uid := users.AllocUID()
	if uid == 0 {
		t.Fatalf("Failed to alloc user")
	}
	u, err := users.Get(uid)
	if err != nil {
		t.Fatal(err)
	}
	u.Password = "123456"
	u, err = users.Get(uid)

	MESSAGE := u.GetMESSAGE()
	if MESSAGE != nil {
		t.Fatal("Failed to get empty message")
	}

	u.AddMESSAGE(types.Message_t{
		Source:  0,
		Content: "test",
		Type:    users.SystemAnnouncment,
	})
	MESSAGE = u.GetMESSAGE()
	if MESSAGE == nil {
		t.Fatal("Failed to get or add message")
	}

}

func TestUserAddFriendAndSendMessage(t *testing.T) {
	//tesing users adding each other as friends and sending message to each other

	uid := users.AllocUID()
	if uid == 0 {
		t.Fatalf("Failed to alloc user")
	}
	ua, err := users.Get(uid)
	if err != nil {
		t.Fatal(err)
	}
	ua.Password = "123456"
	ua, err = users.Get(uid)

	uidb := users.AllocUID()
	if uid == 0 {
		t.Fatalf("Failed to alloc user")
	}
	ub, err := users.Get(uidb)
	if err != nil {
		t.Fatal(err)
	}
	ub.Password = "123456"
	ub, err = users.Get(uidb)

	ua.AcceptAsFriend(ub.Id)
	ub.AcceptAsFriend(ua.Id)

	if ua.IsFriend(ub.Id) == false {
		t.Fatal("failed to add friend")
	}
	if ub.IsFriend(ua.Id) == false {
		t.Fatal("failed to add friend")
	}

	ua.SendMessage(ub.Id, "PING")
	MESSAGE := ub.GetMESSAGE()
	m := types.Message_t{
		Content: "PING",
		Type:    users.UserMessage,
		Source:  ua.Id,
	}
	if MESSAGE == nil || m != *MESSAGE {
		t.Fatal("Wrong message")
	}
}
