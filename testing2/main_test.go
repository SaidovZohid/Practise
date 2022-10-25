package testing2

import "testing"

func TestXxx(t *testing.T) {
	TestCases := []struct{
		input int
		output int
	}{
		{12, 21},
		{121, 121},
		{12344321, 12344321},
		{90, 9},
		{454, 454},
		{3432, 2343},
		{56, 65},
		{74, 47}, 
	}
	for index, testCase := range TestCases {
		result := IsPalindromeNumber(testCase.input)
		if result == testCase.output {
			t.Logf("The Test Case -> %v is PASSED\n", testCase)
		} else {
			t.Errorf("The Test Case -> %v is FAILED, Expected %v, Got %v\n", testCase, testCase.output, result)
			t.Fatalf("ERROR in %d\n", index + 1)
		}
	}
}