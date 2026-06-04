func isAnagram(s string, t string) bool {
	var freq [26]int
	if len(s)!=len(t){
		return false
	}
	for i:=0; i<len(s); i++{
		freq[s[i]-'a']++
		freq[t[i]-'a']--
	}

	for i:=0; i<26; i++{
		if freq[i] !=0{
			return false
		}
	}

	return true
}
