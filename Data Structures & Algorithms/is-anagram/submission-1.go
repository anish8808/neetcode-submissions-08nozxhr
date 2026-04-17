func sortStrings(str string ) string {
	r := []rune(str)

	sort.Slice(r , func(i , j int) bool{
		return r[i]<r[j]
	})

	return string(r)
}

func isAnagram(s string, t string) bool {
	s = sortStrings(s)
	t = sortStrings(t)

	if s == t {
		return true
	}

	return false
}
