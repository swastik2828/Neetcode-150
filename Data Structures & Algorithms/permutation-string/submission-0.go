func checkInclusion(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if len(s1) > len(s2) {
		return false
	}

	var win, need [26] int
	for i := 0; i < n; i++ {
		need[s1[i] - 'a']++
		win[s2[i] - 'a']++
	}

	if need == win {
		return true
	}

	for right := n; right < m; right++ {
		win[s2[right]-'a']++
		win[s2[right-n]-'a']--
		if need == win {
			return true
		}
	}

	return false


}
