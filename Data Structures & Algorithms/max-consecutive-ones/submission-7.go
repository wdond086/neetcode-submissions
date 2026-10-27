func findMaxConsecutiveOnes(nums []int) int {
	maxCount, currentCount := 0, 0
    for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 0
		}
	}
	return maxCount
}
