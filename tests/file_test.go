package tests

import (
	"bytes"
	"mentorChatBackend/models/files"
	"mentorChatBackend/models/types"
	"testing"
)

func TestFile(t *testing.T) {
	buffer := bytes.NewBufferString("TESTTESTTEST")
	fileid := files.NewFile(buffer.Bytes())
	if fileid == "" {
		t.Fatal("Failed to create file ID")
	}
	data, err := files.GetFile(fileid)
	if err != nil {
		t.Fatal(err)
	} else if string(data) != "TESTTESTTEST" {
		t.Fatal("Failed to retrieve file but got : ", string(data))
	}
}
