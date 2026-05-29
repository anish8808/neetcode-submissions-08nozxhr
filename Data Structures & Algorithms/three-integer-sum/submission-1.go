func threeSum(nums []int) [][]int {
	result := [][]int{}
	sort.Slice(nums , func (a , b int)bool {
		return nums[a]<nums[b]
	})

	

	for i:=0; i<len(nums)-2; i++{
	
		if i>0 && nums[i]==nums[i-1]{
			continue	
		}

		left :=i+1
		right := len(nums)-1

		for left <right{
			
			sum := nums[left]+ nums[right] + nums[i]

			if sum ==0 {
				result = append (result , []int{nums[i] , nums[left] , nums[right]})

				left++
				right --

				for left<right && nums[left]==nums[left-1]{
					left++
				}

				for left<right && nums[right]==nums[right+1]{
					right--
				}

			} else if sum<0{
				left++
			}else {
				right --
			}
		}
	}

	return result
}
