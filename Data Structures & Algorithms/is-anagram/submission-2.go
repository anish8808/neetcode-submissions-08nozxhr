func isAnagram(s string, t string) bool {
	if len(s) != len(t){
		return false
	}

	countArray := [26]int{}

	for i:=0; i<len(s); i++{
		countArray[s[i]-'a']++
		countArray[t[i]-'a']--
	}

	for _ , val := range countArray {
		if val>0{
			return false
		}
	}

	return true
}
