class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        maxCount = 0
        currentMax = 0
        for i in range(len(nums)):
            if nums[i] == 1:
                currentMax += 1
            else:
                if currentMax > maxCount:
                    maxCount = currentMax
                currentMax = 0
            if i == len(nums) - 1:
                return max(currentMax, maxCount)