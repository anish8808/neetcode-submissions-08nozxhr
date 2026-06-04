func hasDuplicate(nums []int) bool {
    freqMap := make(map[int]int)

	for i:=0; i<len(nums); i++{
		if freqMap[nums[i]]>=1{
			return true
		}
		freqMap[nums[i]]++
	}

	return false
}
