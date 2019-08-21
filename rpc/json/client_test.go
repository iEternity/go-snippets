package gob

import "testing"

func TestClient(t *testing.T) {
	client, err := DialEchoService("tcp", "10.246.44.205:8888")
	if err != nil {
		t.Fatal("DialEchoService error:", err)
	}

	msg := Request{
		Name:  "jack",
		Age:   25,
		Phone: []string{"1234567", "9887434", "34567"},
		Data:  map[string]string{"addr": "abcd", "city": "beijing"},
	}

	var reply Reply

	if err := client.Echo(msg, &reply); err != nil {
		t.Fatal("Echo error:", err)
	}

	t.Logf("msg: %v, reply: %v", msg, reply)
}
