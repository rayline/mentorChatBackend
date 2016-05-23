package files

import (
	"io/ioutil"
	"mentorChatBackend/models/types"
	"strconv"
)

func GetFile(Id types.FileID_t) (data []byte, err error) {
	data, err = ioutil.ReadFile("static/userfiles/" + Id)
	return data, err
}

func GetOwnerOfFile(Id types.FileID_t) (data types.UserID_t, err error) {

}
