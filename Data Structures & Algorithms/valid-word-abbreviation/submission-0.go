func validWordAbbreviation(word string, abbr string) bool {
	n,m := len(word), len(abbr)
	i, j := 0, 0
	for i < n && j < m {
		if abbr[j] == '0' {
			return false
		}

		if abbr[j] >= 'a' && abbr[j] <= 'z' {
			if word[i] == abbr[j] {
				i++
				j++
			} else {
				return false
			}
		} else {
			subLen := 0
			for j < m && abbr[j] >= '0' && abbr[j] <= '9' {
				subLen = subLen * 10 + int(abbr[j] - '0')
				j++
			}
			i += subLen
		}
	}
	return i == n && j == m
}
