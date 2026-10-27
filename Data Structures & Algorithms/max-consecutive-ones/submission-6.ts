class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    findMaxConsecutiveOnes(nums: number[]): number {
        let currentCount = 0;
        let maxCount = 0;
        for (let i = 0; i < nums.length; i++) {
            currentCount = nums[i] === 1 ? currentCount + 1 : 0;
            if (currentCount > maxCount) maxCount = currentCount;
        }
        return maxCount;
    }
}
