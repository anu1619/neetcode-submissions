func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	freq := make([][]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}
	for ky, v := range count {
		freq[v] = append(freq[v], ky)
	}
	arr := make([]int, 0)
	for i := len(freq) - 1; i >= 0; i-- {

		if len(freq[i]) != 0 {
			arr = append(arr, freq[i]...)
		}
		if len(arr) == k {
			break
		}
	}
	return arr
}