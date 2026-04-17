func abs(x ,y int) int {
	val := x- y
	fmt.Println(val)
	if val<0{
		return -1 * val
	}

	return val
}

func containsNearbyDuplicate(nums []int, k int) bool {
	countMap := make(map[int]int )

	for i:=0; i<len(nums); i++{
		if countMap[nums[i]] >0{
			val := abs(i+1 , countMap[nums[i]])

			if val <=k	{
				return true
			}
		}

		countMap[nums[i]] = i+1
	}

	return false;
}
