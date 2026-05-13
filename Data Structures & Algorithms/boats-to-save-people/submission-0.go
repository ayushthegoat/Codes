func numRescueBoats(arr []int, limit int) int {
    sort.Ints(arr)
	left := 0
	right := len(arr) - 1
	boats := 0

	for left <= right {
		if arr[left] + arr[right] <= limit {
			left++
			right--
			boats++
		} else {
			boats++
			right--
		} 

	}
	return boats
}
