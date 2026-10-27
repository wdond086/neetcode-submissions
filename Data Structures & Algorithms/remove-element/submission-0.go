func removeElement(nums []int, val int) int {
	length := len(nums)
	// Start from the right and move towards the left as you are moving things towards the right as you delete
	for i := length - 1; i >= 0; i-- {
		if nums[i] == val {
			nums[i] = 0
			for j := i + 1; j < length; j++ {
				nums[j-1] = nums[j]
			}
			length--
		}
	}
	return length
}
