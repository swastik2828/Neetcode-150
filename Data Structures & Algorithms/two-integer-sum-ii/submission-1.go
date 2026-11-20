import "slices"
func twoSum(nums []int, target int) []int {
	left := 0
	right := len(nums)-1
	slices.Sort(nums)

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}

	}
	return nil
}
