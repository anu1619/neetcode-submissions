func isPalindrome(s string) bool {
left := 0
right := len(s)-1
for left < right{
	if !isAlphaNumeric(s[left]){
		left++
		continue
	}
	if !isAlphaNumeric(s[right]){
		right--
		continue
	}
	if lower(s[left]) != lower(s[right]){
		return false
	}
	left++
	right--
}
return true
}

func lower(c byte)byte{
	if c >= 'A' && c <= 'Z'{
		return c+('a'-'A')
	}
	return c
}

func isAlphaNumeric(c byte)bool{
	if ('A' <= c && c<='Z') ||('a'<= c && c<= 'z') || ('0'<=c && c<='9'){
		return true
	}
	return false
}
