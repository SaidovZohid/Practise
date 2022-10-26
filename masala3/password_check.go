package masala3

import (
	u "unicode"
)

func CheckPassword(pswrd string) bool {
	if len(pswrd) < 8 {
		return false
	} 
	var symbCount, lowerCount, upperCount, numCount, alphaCount int
	for _, v := range pswrd {
		if u.IsDigit(v) {
			numCount++
			continue
		}
		if u.IsLetter(v) && u.IsLower(v){
			lowerCount++
			alphaCount++
			continue
		} else if u.IsLetter(v) && u.IsUpper(v){
			upperCount++
			alphaCount++
			continue
		}
		if u.IsPunct(v){
			symbCount++
			continue
		}
	}
	if numCount > 0 && symbCount > 0 && lowerCount > 0 && upperCount > 0 && alphaCount > 0 {
		return true
	} else {
		return false
	}
}