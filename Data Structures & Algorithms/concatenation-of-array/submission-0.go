func getConcatenation(nums []int) []int {
	numsLength := len(nums)
	ans := make([]int, numsLength * 2)
    for i := range nums {
		ans[i] = nums[i]
		ans[i + numsLength] = nums[i]
	}
	return ans
}
