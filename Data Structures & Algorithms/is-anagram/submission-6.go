func isAnagram(s string, t string) bool {
	sbyte := []byte(s)
	tbyte := []byte(t)
	if len(s)!=len(t){
		return false
	}

	sort.Slice(sbyte , func(a , b int)bool{
		return sbyte[a] <sbyte[b]
	})

	sort.Slice(tbyte , func(a , b int)bool{
		return tbyte[a] <tbyte[b]
	})
	for i:=0; i<len(s); i++{
		if sbyte[i] != tbyte[i]{
			return false
		}
	}

	return true
}
