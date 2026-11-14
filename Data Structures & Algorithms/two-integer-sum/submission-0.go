func twoSum(nums []int, target int) []int {
    // n := len(nums)
	mpp := make(map[int]int)

	for i, num := range nums {
		a := num
		more := target - a
		if val, ok := mpp[more]; ok {
			return []int{val, i}
		}
		mpp[a] = i
	}
	return []int{-1,-1}
}
