func hasDuplicate(nums []int) bool {
	occurrence := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		if _, exists := occurrence[nums[i]]; exists {
			return true
		}
		occurrence[nums[i]] = struct{}{}
	}
	return false
}
