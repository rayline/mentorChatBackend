package types

type UserID_t uint64
type TokenID_t uint64
type Password_t string
type FileID_t string
type Message_t struct {
	Source  UserID_t
	Type    string
	Content string
}
