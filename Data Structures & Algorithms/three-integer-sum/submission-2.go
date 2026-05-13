func threeSum(arr []int) [][]int {
    res := [][]int{}
	sort.Ints(arr)

	for i:= 0; i<= len(arr) - 2; i++ {
		//we are skipping the duplicated i here becase 
		//if not we wil get duplicate anser 
		// also check i > 0 becoz becasue we must not get array out of bound 
		if i > 0 && arr[i] == arr[i - 1] {
			continue
		}

		//we fix the first i in the aboeve loop and the n
		//we satrt the lop for teh rest of the array
		left, right := i+1, len(arr) - 1
		for left < right {
			sum := arr[i] + arr[left] + arr[right]
			if sum  == 0 {
				res = append(res, []int{arr[i], arr[left], arr[right]})
				//our result is appended
				//now we check and remove duplciates
				for left < right && arr[left] == arr[left+1] {
					left++
					//moving left forward until we skip all the dupes
				}
				for left < right && arr[right] == arr[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0{
                right--
			} else {
				left++
			}
		}
	}
	return res
}
