func isValid(s string) bool {
	var st []byte

	for i:=0; i<len(s); i++{
		if s[i] == '(' || s[i] =='{' || s[i] =='['{
			st = append(st , s[i])
		}else {

			if len(st)==0{
				return false
			}

			top := st[len(st)-1]
			if (s[i] == ')' && top == '(') || (s[i] == '}' && top == '{') || (s[i] == ']' && top == '['){
				st = st[:len(st)-1]
			}else {
				return false 
			}
		}
	}

	if (len(st) != 0){
		return false
	}


	return true
}
