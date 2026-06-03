func max(a , b int)int {
	if a>b{
		return a
	}

	return b
}

func characterReplacement(s string, k int) int {
left :=0
longest :=0
freqMap := make(map[byte]int)
maxFreq :=0
for right :=0; right<len(s); right++{
	freqMap[s[right]]++

	maxFreq = max(maxFreq , freqMap[s[right]])

	for (right-left+1)-maxFreq >k{
		freqMap[s[left]]--
		left++
	}

	longest = max(longest , right-left+1)
}	

return longest
}
