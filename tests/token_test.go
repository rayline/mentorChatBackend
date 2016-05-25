package tests

import "testing"
import "mentorChatBackend/models/tokens"
import "mentorChatBackend/models/users"

func TestTokens(t testing.T) {
	//testing basic token creating and retrieving
	uid := users.AllocUID()
	if uid == 0 {
		t.Fatalf("Failed to alloc user")
	}
	u, err := users.Get(uid)
	if err != nil {
		t.Fatal(err)
	}
	u.Password = "123456"
	users.Set(uid, *u)
	u, err = users.Get(uid)
	if err != nil {
		t.Fatal(err)
	}
	if u.Password != "123456" {
		t.Fatal("Failed to change password")
	}

	token := tokens.NewToken(uid)
	uid2, err := tokens.Get(token)
	if err != nil {
		t.Fatal(err)
	}
	if uid2 != uid {
		t.Fatal("Failed to retrieve token")
	}
}
