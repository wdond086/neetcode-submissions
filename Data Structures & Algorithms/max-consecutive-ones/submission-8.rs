impl Solution {
    pub fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
        let mut best : i32 = 0;
        let mut current : i32 = 0;
        for v in nums {
            if v == 1 {
                current += 1;
                if current > best {
                    best = current;
                }
            } else {
                current = 0;
            }
        }
        return best;
    }
}
