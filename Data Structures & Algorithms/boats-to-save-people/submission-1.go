func numRescueBoats(arr []int, limit int) int {
    sort.Ints(arr)
	left := 0
	right := len(arr) - 1
	boats := 0

	 for left <= right {
        if arr[left]+arr[right] <= limit {
            // Case 1: Both fit in one boat
            left++
            right--
            boats++
        } else if arr[left]+arr[right] > limit {
            // Case 2: Sum exceeds limit - heavy goes alone
            right--
            boats++
        } else {
            // Case 3: The remaining person (when left == right)
            boats++
            break
        }
    }
    return boats
	
}
