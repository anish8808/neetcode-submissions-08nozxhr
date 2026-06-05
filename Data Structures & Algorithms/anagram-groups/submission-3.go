func sortString(s string) string{
	r := []rune(s)

	sort.Slice(r , func (a , b int)bool{
		return r[a]<r[b]
	})

	return string(r)
}

func groupAnagrams(strs []string) [][]string {
	freqMap := make(map[string][]string)

	for i:=0; i<len(strs); i++{
		sortedString := sortString(strs[i])

		freqMap[sortedString] = append(freqMap[sortedString] , strs[i])
	}
	var ans [][]string
	for _ , val := range freqMap{
		ans  = append(ans ,  val)
	}

	return ans
}
