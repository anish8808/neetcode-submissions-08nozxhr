func productExceptSelf(nums []int) []int {
	size := len(nums)

	suffix := make ([]int , size)
	prefix := make([]int , size)

	prefix[0]=1
	suffix[size-1] = 1

	for i :=1; i<len(nums); i++{
		prefix[i] = prefix[i-1] * nums[i-1]
	}

	for i := size-2; i>=0; i-- {
		suffix[i] = suffix[i+1]*nums[i+1]
	}

	final := make([]int ,size)

	for i:=0; i<len(nums); i++{
		final[i] = prefix[i]*suffix[i]
	}

	return final
}
