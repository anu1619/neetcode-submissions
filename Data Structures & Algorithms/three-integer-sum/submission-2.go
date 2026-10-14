func threeSum(nums []int) [][]int {
output := [][]int{}
entryMap := map[[3]int]struct{}{}
 nums = mergeSort(nums)
 for i:=0;i<len(nums);i++{
	left := i+1
	right := len(nums)-1
	for left<right{
	if -nums[i] == nums[left]+nums[right]{
		if _, ok := entryMap[[3]int{nums[i], nums[left], nums[right]}]; !ok{
		output = append(output, []int{nums[i], nums[left], nums[right]})
		entryMap[[3]int{nums[i], nums[left], nums[right]}] = struct{}{}
		}
	}
	if nums[left]+nums[right] < -nums[i]{
		left++
	}else {
		right--
	}
	}
 }
 return output
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
