func isAnagram(s string, t string) bool {
		if len(s) != len(t) {
			return false
		}

		charArry := [26]int{}

		for i:=0; i<len(s); i++{
			charArry[s[i]-'a']++
			charArry[t[i]-'a']--
		}

		for _ , val := range charArry{
			if val != 0{
				return false 
			}
		}

		return true
}
