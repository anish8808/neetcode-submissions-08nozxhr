func isPalindrome(s string , start , end int)bool{
    for start<end{
        if s[start]==s[end]{
            start++
            end--
        }else {
            return false
        }
    }

    return true
}

func validPalindrome(s string) bool {
    start :=0
    end := len(s)-1

    for start<end{
        if s[start] == s[end]{
            start++
            end--
        }else {
            left := isPalindrome(s , start+1 , end)
            right := isPalindrome(s , start , end-1)

            return left || right
        }
    }


    return true
}
