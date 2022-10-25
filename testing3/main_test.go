package testing3

import "testing"

func TestIsPalindrome(t *testing.T) {
	TestCases := []struct{
		input string
		output bool
	}{
		{"aba", true},
		{"wallet", false},
		{"karima", false},
		{"hello", false},
		{"oppo", true},
		{"komil", false},
		{"ehtimol", false},
		{"aziza", true},
		{"fef", true},
		{"teta", false},
		{"va", false},
		{"adda", true},
	}
	for _, testcase := range TestCases {
		if res := isPalindrome(testcase.input); res != testcase.output {
			t.Errorf("The test case -> %v, Expected %v, Got %v\n", testcase, testcase.output, res)
		} else {
			t.Log("PASSED")
		}
	}
}