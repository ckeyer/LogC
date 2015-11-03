package main

import (
	"bytes"
	"net/http"
	"testing"
)

// TestPush ...
func TestPush(t *testing.T) {
	send("hellsdf")
	send("hlknxcvinfd")
	code := send("hello world")
	if code != http.StatusCreated {
		t.Error("no done...")
	}

}

func send(data string) int {
	r := bytes.NewReader([]byte(data))
	re, _ := http.Post("http://127.0.0.1:8087", "application/json", r)
	return re.StatusCode
}
