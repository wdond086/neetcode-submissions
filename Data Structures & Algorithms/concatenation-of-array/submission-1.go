func getConcatenation(nums []int) []int {
    for i := range nums {
		nums = append(nums, nums[i])
	}
	return nums
}
