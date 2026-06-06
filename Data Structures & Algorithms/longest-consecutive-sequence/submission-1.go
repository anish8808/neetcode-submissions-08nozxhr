func longestConsecutive(nums []int) int {
	longest := 0
	

	freqMap := make(map[int]int)

	for i:=0; i<len(nums); i++{
		freqMap[nums[i]]++
	}

	for key , _ := range freqMap{
		
		count := 0
		_ , ok := freqMap[key-1]
		if !ok {
			curr := key
			count =1

			check := true

			for check {
				_ , ok := freqMap[curr+1]
				if ok{
					count++
					curr ++
				}else{
					check = false
				}
			}
		}

		longest = max(longest ,count )
	}

	return longest
}
