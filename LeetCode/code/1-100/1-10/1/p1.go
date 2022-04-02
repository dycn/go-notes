package p1

func twoSum(nums []int, target int) []int {
	s := make(map[int]int, len(nums))
	for k, v := range nums {
		k1, ok := s[v]
		if ok {
			return []int{k1, k}
		} else {
			s[target-v] = k
		}
	}
	return []int{}
}
