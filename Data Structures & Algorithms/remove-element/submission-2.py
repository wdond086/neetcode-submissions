class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        length = len(nums)
        for i in range(length - 1, -1, -1):
            if nums[i] == val:
                for j in range(i + 1, length):
                    nums[j - 1] = nums[j]
                nums[length - 1] = 0
                length -= 1
        return length
        