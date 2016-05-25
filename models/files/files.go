package files

import (
	"io/ioutil"
	"log"
	"math/rand"
	"mentorChatBackend/models/types"
	"os"
	"strconv"
)

func GetFile(Id types.FileID_t) (data []byte, err error) {
	data, err = ioutil.ReadFile("static/userfiles/" + string(Id))
	return data, err
}

func NewFile(data []byte) types.FileID_t {
	for Id := types.FileID_t(""); Id == ""; {
		Id = types.FileID_t(types.FileID_t(randomString()))
		os.Mkdir("static/userfiles/", os.ModePerm)
		f, err := os.OpenFile("static/userfiles/"+string(Id), os.O_CREATE|os.O_EXCL, os.ModePerm)
		if err != nil {
			log.Println(err)
			Id = ""
		} else {
			f.Write(data)
			return Id
		}
	}
	return ""
}

func randomString() string {
	a := rand.Int63()
	b := rand.Int63()
	return strconv.FormatInt(a, 16) + strconv.FormatInt(b, 16)
}
