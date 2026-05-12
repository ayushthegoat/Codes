func removeDuplicates(arr []int) int {
    n := len(arr)
    i := 0
	j := 1
	for j < n {
		if arr[j] == arr[i] {
			j++
           continue
		} else if arr[j] != arr[i] {
		   i++	
           arr[i] = arr[j]
		   j++
		} 
        
	}
	return i + 1
}
