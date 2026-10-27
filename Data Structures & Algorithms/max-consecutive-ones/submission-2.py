class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        maxCount = 0
        currentCount = 0
        for i in range(len(nums)):
            currentCount = currentCount + 1 if nums[i] == 1 else 0
            maxCount = max(currentCount, maxCount)
        return maxCount
        