func max(x , y int) int{
	if x>y{
		return x
	}

	return y
}

func characterReplacement(s string, k int) int {
	freqMap := make(map[byte]int)
	longest :=0
	maxfreq := 0
	left :=0
	for right :=0; right <len(s); right++{
		freqMap[s[right]]++

		maxfreq = max(maxfreq , freqMap[s[right]])
		for((right-left+1)-maxfreq>k){
			freqMap[s[left]]--
			left++
		}

		longest = max(longest , right-left+1)
	}

	return longest
}
