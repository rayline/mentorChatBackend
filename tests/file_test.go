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
	data := files.GetFile(fileid)
	if string(data) != "TESTTESTTEST" {
		t.Fatal("Failed to retrieve file but got : ", string(data))
	}
}
