package masala

import (
	"testing"
)

func TestTriangle(t *testing.T) {
	TestCases := []struct{
		a int
		b int
		c int
		output bool
	}{
		{6, 8, 10, true},
		{3, 9, 42, false},
	}
	for _, testcase := range TestCases {
		if res := FindTriangle(testcase.a, testcase.b, testcase.c); res != testcase.output {
			t.Errorf("The Test Case -> %v, Expected %v, Got %v\n", testcase, testcase.output, res)
		} else {
			t.Log("PASSED")
		}
	}
}