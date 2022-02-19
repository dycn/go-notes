package p1

func twoSum(nums []int, target int) []int {
    s := make(map[int]int, len(nums))
    for i := 0; i < len(nums); i++ {
        value := nums[i]
        index, ok := s[value]
        if ok {
            return []int{index, i}
        } else {
            s[target-value] = i
        }
    }
    return []int{}
}