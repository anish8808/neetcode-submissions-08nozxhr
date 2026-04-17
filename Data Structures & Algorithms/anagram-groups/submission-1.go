func sortString(str string ) string {
	r := []rune(str)

	sort.Slice(r , func(i  , j int)bool {
		return r[i] < r[j]
	})

	return string(r)
}
func groupAnagrams(strs []string) [][]string {
	freqMap := make(map[string][]string)

	for i:=0; i<len(strs); i++{
		s := sortString(strs[i])

		freqMap[s] = append (freqMap[s] , strs[i])
	}

	ans := [][]string{}

	for _ , val := range freqMap{
		ans = append(ans , val)
	}


	return ans
}
