func twoSum(nums []int, target int) []int {
    freqMap := make(map[int]int)

	for i:=0; i<len(nums); i++{


		val , ok := freqMap[target - nums[i]]

		if ok {
			ans := []int{val , i}
			return ans
		}

		freqMap[nums[i]]= i 
	}

	return []int{}
}
