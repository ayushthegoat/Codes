func validPalindrome(s string) bool {
  i := 0
  j := len(s) - 1
  for i < j {
     if s[i] != s[j] {
		//agar nahi hai to ONLY EK HI CHANCE HAI ?
		//YA to left to hatao tab banega 
		//YA for right to hatao tab banega palli
		//nahi bana to out 
		//so we send the isPalli fucntion with a bound change and check
		//it handles everything else 
		return isPalli(s, i+1, j) || isPalli(s, i, j-1) 
	 }
	 i++
	 j--
  }
  return true
}

func isPalli(s string, left, right int) bool {
	i:=left
	j:=right
	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}
