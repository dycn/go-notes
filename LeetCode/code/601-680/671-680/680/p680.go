package p680

import (
	"fmt"
	"strings"
)

func ValidPalindrome(s string) bool {
	for k, _ := range s {
		newStr := append([]string{}, s[:k], s[k+1:])
		newStr1 := strings.Join(newStr, "")
		if isHuiwen(newStr1) {
			return true
		}
	}
	return false
}

func isHuiwen(s string) bool {
	fmt.Println(s)
	var flag bool = true
	if s == "" {
		return flag
	}
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			flag = false
			break
		}
	}
	return flag
}
