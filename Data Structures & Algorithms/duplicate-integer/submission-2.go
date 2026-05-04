func hasDuplicate(nums []int) bool {
    freqMap := make(map[int]bool)

	for _ , val := range nums {
		fmt.Println(val)
		if (freqMap[val]){
			return true
		}

		freqMap[val] = true
	}
	return false
}
