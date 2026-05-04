func sortString(str string) string{
	b := []rune(str)
	sort.Slice(b , func(i , j int)bool {
		return b[i]<b[j]
	})	

	return string(b)
}

func isAnagram(s string, t string) bool {
	if (len(s)!= len(t)){
		return false 
	}

	s = sortString(s)
	t = sortString(t)

	for i:=0; i<len(s); i++{
		if s[i] != t[i]{
			return false
		}
	}

	return true

}
