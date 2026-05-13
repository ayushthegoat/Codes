func twoSum(arr []int, target int) []int {
  n := len(arr)
  l := 0
  r := n - 1

  for l < r {
	if arr[l] + arr[r] == target {
		return []int{l+1, r+1}
	} 
	if arr[l] + arr[r] > target {
		r--
	} else {
		l++
	}
  }
  return []int{0,0}
}
