func threeSum(nums []int) [][]int {
    res := [][]int{}
	sort.Ints(nums)

	for i:= 0; i< len(nums); i++ {
	    curr:= nums[i]
		if curr > 0 {
			break
		}
		if i > 0 && curr == nums[i-1] {
			continue
		}
		l,r:= i+1, len(nums)-1
		for l < r{
			sum:= curr + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				res = append(res, []int{curr, nums[l], nums[r]})
				l++
				r--

			    for l < r && nums[l] == nums[l - 1] {
					l++
				}
			}
		}
	}
	return res
}
