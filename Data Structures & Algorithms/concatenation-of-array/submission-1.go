func getConcatenation(nums []int) []int {
    var ans []int

    for _ , val := range nums {
        ans = append(ans , val)
    }

    for _ , val := range nums {
        ans = append(ans , val)
    }

    return ans

}
