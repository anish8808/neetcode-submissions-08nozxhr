func majorityElement(nums []int) int {
    freqMap := make(map[int]int)

	for i:=0; i<len(nums); i++{
		freqMap[nums[i]]++
	}


	for val , count := range freqMap{
		if count>(len(nums)/2){
			return val
		}
	}


	return -1
}
