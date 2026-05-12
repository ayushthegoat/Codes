func removeDuplicates(arr []int) int {
    n := len(arr)
    i := 0
	j := 1
	for j < n {
		if arr[i] != arr[j] {
			i++
			arr[i] = arr[j]
			
		}
		j++
        
	}
	return i + 1
}
