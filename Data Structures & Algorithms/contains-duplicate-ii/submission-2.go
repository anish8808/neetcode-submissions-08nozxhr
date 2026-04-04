func containsNearbyDuplicate(nums []int, k int) bool {
    count_map := make(map[int]int)

    for i :=0; i<len(nums); i++{

        _ , ok := count_map[nums[i]]

        if ok {
            if i - count_map[nums[i]] <=k {
                return true
            }
        }

        count_map[nums[i]] = i
    }

    return false
}
