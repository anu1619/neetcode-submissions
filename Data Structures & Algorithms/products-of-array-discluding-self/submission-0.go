func productExceptSelf(nums []int) []int {
	arr := make([]int, len(nums))
	prefix := 1
	postfix:= 1
	for i, v := range nums{
		arr[i] = prefix
		prefix *= v
	}
	for i:=len(nums)-1;i>=0;i--{
		arr[i] *= postfix
		postfix *= nums[i]
	}
	return arr
}

