func isPalindrome(s string) bool {
reg, _ := regexp.Compile("[^a-zA-Z0-9]+")
s = reg.ReplaceAllString(s,"")
s = strings.ToLower(s)
left := 0
right := len(s)-1
for left < right{
	if s[left] != s[right]{
		return false
	}
	left++
	right--
}
return true
}
