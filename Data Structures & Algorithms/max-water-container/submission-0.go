func maxArea(heights []int) int {
	res,l := 0,0
	n := len(heights)
	r := n-1

	for l < r && l < n-1 && r > 0 {
		area := (r-l) * min(heights[l], heights[r])
		res = max(res, area)
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
