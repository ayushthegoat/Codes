func rotate(nums []int, k int) {
   k = k % len(nums)
   if k < 0 {
	k = k + len(nums)
   }
   //part 1 reverser
   reverse(nums, 0, len(nums) - k - 1)
   //part 2 reverse
   reverse(nums, len(nums) - k, len(nums) - 1)
   //full reverse
   reverse(nums, 0, len(nums) - 1)
}
func reverse(nums []int, front, rear int) {
	i:=front
	j:=rear

	for i < j {
       nums[i], nums[j] = nums[j], nums[i]
	   i++
	   j--
	}
}


