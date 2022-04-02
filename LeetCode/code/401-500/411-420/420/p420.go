package p420

func strongPasswordChecker(password string) int {
	checkRepeat(password)
}

func checkRepeat(s string) int {
	var count int
	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] && s[i-1] == s[i] {
			count++
			i += 2
		}
	}
	return count
}
