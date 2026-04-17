func hasDuplicate(nums []int) bool {
    freqMap := make(map[int]bool)

	for i:=0; i<len(nums); i++{
		if freqMap[nums[i]] {
			return true
		}
		freqMap[nums[i]] = true
	}

	return false

}
