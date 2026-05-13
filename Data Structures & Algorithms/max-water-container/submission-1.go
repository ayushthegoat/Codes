func maxArea(arr []int) int {
  left:= 0
  right:= len(arr)-1
  maxCap := 0
  for left < right {
	capN:= (right - left) * min(arr[left], arr[right])
    if capN > maxCap {
		maxCap = capN
	}
	if arr[left] > arr[right] {
		right--
	}else {
		left++
	}
  }
  return maxCap
}

