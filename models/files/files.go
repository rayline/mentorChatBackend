package files

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"math/rand"
	"mentorChatBackend/models/types"
	"os"
	"strconv"
)

func GetFile(Id types.FileID_t) (data []byte, err error) {
	data, err = ioutil.ReadFile("static/userfiles/" + string(Id))
	return data, err
}

func NewFile(data []byte) (types.FileID_t, error) {
	hash := sha256.Sum256(data)
	Id := types.FileID_t(hex.EncodeToString(hash[0:32]))
	err := ioutil.WriteFile("static/userfiles/"+string(Id), data, os.ModePerm)
	return Id, err
}

func randomString() string {
	a := rand.Int63()
	b := rand.Int63()
	return strconv.FormatInt(a, 16) + strconv.FormatInt(b, 16)
}
