func moveZeroes(arr []int) {
	i:= 0
	j:= 0

	for j < len(arr) {
		
		if arr[j] > 0 {
			// arr[i] = arr[j]
			// arr[j] = 0 
			arr[i],arr[j] = arr[j],arr[i]
			i++
		}
		j++
	}
	
}
