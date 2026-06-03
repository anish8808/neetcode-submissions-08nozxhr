func max(a ,b int) int {
	if a>b{
		return a
	}

	return b
}

func lengthOfLongestSubstring(s string) int {
	left :=0
	longest :=0
	freqMap := make(map[byte]int)
	for right:=0; right<len(s); right++{
		freqMap[s[right]]++

		for freqMap[s[right]]>1{
			freqMap[s[left]]--
			left++
		} 

		longest = max (longest , right-left +1)
	}

	return longest
}
