func isValid(s string) bool {
    var stack []byte 

	for c :=0; c<len(s); c++{
		if s[c]=='(' || s[c] =='{' || s[c] == '[' {
			stack = append(stack , s[c])
		} else {

		if len(stack) ==0{
			return false 
		}
		top := stack[len(stack)-1]
		if (s[c] ==')' && top =='(') || (s[c] =='}' && top =='{') || (s[c] ==']' && top =='['){
			stack = stack[:len(stack)-1]
		}else {
			return false
		}
		}
	}

	if len(stack) > 0{
		return false
	}

	return true
}
