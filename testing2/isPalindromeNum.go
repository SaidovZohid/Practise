package testing2

func IsPalindromeNumber(num int) int {
	var res int
	for num > 0 {
		res = res * 10 + num % 10
		num /= 10
	}
	return res
}