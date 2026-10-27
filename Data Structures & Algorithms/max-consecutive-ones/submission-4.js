class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    findMaxConsecutiveOnes(nums) {
        let currentCount = 0;
        let maxCount = 0;
        for (let i of nums) {
            currentCount = (i == 1) ? currentCount + 1 : 0;
            maxCount = Math.max(maxCount, currentCount);
        }
        return maxCount;
    }
}
