func sortString(str string)string{
	b := []byte(str)

	sort.Slice(b , func(a , c int)bool{
		return b[a]<b[c]
	})

	return string(b)
}

func groupAnagrams(strs []string) [][]string {
 freqmap := make(map[string][]string)

 for i:=0; i<len(strs); i++{
	sortStr := sortString(strs[i])

	freqmap[sortStr] = append(freqmap[sortStr] , strs[i])
 }

	var ans [][]string
	for _ , val := range freqmap{
		ans = append(ans , val)
	}

	return ans
}
