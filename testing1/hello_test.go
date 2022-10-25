package testing1

import "testing"

func TestHello(t *testing.T) {
	res := Hello("Hello ", "Zohid")
	if res == "Hello Zohid" {
		t.Logf("%s\n", res)	
	} else {
		t.Errorf("FAILED. Becaue string -> %s\n", res)
	}
}