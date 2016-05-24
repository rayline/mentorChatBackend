package files

import (
	"io/ioutil"
	"math/rand"
	"mentorChatBackend/models/types"
	"os"
	"strconv"
)

func GetFile(Id types.FileID_t) (data []byte, err error) {
	data, err = ioutil.ReadFile("static/userfiles/" + Id)
	return data, err
}

func NewFile(data []byte) types.FileID_t {
	for Id := types.FileID_t(""); Id == ""; {
		Id = types.FileID_t(randomString)
		f, err := os.OpenFile("static/userfiles/"+Id, os.O_CREATE|os.O_EXCL, os.ModePerm)
		if err != nil {
			Id = ""
		} else {
			f.Write(data)
			return Id
		}
	}
}

func randomString() string {
	a = rand.Int63()
	b = rand.Int63()
	return strconv.Itoa(a) + strconv.Itoa(b)
}
