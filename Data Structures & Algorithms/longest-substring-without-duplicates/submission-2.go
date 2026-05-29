func max(x ,y int) int{
	
	if x>y {
		return x
	}

	return y
}

func lengthOfLongestSubstring(s string) int {
	left :=0
	longestLen :=0

	freqMap := make(map[byte]int)

	for right :=0; right<len(s); right++{
		freqMap[s[right]]++

		for freqMap[s[right]]>1{
			freqMap[s[left]]--
			left++
		}

		longestLen = max(longestLen , right-left+1)
	
	}

	return longestLen
}
