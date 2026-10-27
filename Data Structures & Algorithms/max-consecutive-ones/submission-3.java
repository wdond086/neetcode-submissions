class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
        int maxCount = 0, currentCount = 0;
        for (int i = 0; i < nums.length; i++) {
            currentCount = (nums[i] == 1) ? currentCount + 1 : 0;
            maxCount = Integer.max(maxCount, currentCount);
        }
        return maxCount;
    }
}