package masala3

import (
	"testing"
)

func TestPassword(t *testing.T) {
	Passswords := []struct{
		pswrd string
		otp bool
	}{
		{"Wew23/Z!r", true},
		{"234der", false},
		{"QertTr/e", false},
		{"Zohid!2004", true},
		{"Zohid2004", false},
	}
	for i, password := range Passswords {
		if res := CheckPassword(password.pswrd); res != password.otp {
			t.Errorf("Test Case %d is FAILED, Expected %v, Got %v\n", i+1, password.otp, res)
			t.Fatal("TEST IS NOT PASSED")
		} else {
			t.Log("PASSED")
		}
	}
}