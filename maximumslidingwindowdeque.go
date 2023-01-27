 func maxSlidingWindow(nums []int, k int) []int {
    n := len(nums)
    if n == 0 || k == 0 {
        return []int{}
    }
    result := make([]int, n-k+1)
    deque := make([]int, 0)
    for i := 0; i < n; i++ {
        for len(deque) > 0 && nums[i] >= nums[deque[len(deque)-1]] {
            deque = deque[:len(deque)-1]
        }
        deque = append(deque, i)
        if i >= k-1 {
            result[i-k+1] = nums[deque[0]]
            if deque[0] == i-k+1 {
                deque = deque[1:]
            }
        }
    }
    return result
}
