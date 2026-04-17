func containsNearbyDuplicate(nums []int, k int) bool {
  countMap := make(map[int]int)

  for i:=0; i<len(nums); i++{
	prevIndex , ok := countMap[nums[i]]
	if ok {
		if i - prevIndex <=k{
			return true
		}
	}

	countMap[nums[i]] = i
  }

  return false
}
