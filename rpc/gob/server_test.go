package gob

import (
	"testing"
)

func TestServer(t *testing.T) {
	err := RunService("tcp", ":8888")
	if err != nil {
		t.Fatal("Listen TCP error:", err)
	}

	select {}
}
