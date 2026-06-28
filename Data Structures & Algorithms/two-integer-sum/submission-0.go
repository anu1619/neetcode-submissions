func twoSum(nums []int, target int) []int {
    remaining := make(map[int]int)
    for i:=0;i<len(nums);i++{
        if v, exists := remaining[nums[i]]; exists{
            return []int{v, i}
        } 
        remaining[target - nums[i]] = i
    }
    return []int{}
}
