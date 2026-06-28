func groupAnagrams(strs []string) [][]string {
	groups := make([][]string, 0)
	marked := make(map[int]struct{})
	for i := 0; i < len(strs); i++ {
        if _, ok := marked[i]; ok {
				continue
			}
		group := []string{strs[i]}
		
		for j := i+1; j < len(strs); j++ {
			count := make(map[byte]int)
			if len(strs[i]) != len(strs[j]) {
				continue
			}
			for k := 0; k < len(strs[i]); k++ {
				count[strs[i][k]]++
				count[strs[j][k]]--
			}
			isAnagram := true
			for _, v := range count {
				if v != 0 {
					isAnagram = false
					break
				}
			}
			if isAnagram {
				group = append(group, strs[j])
				marked[j] = struct{}{}
			}

		}
			groups = append(groups, group)
		
	}
	return groups
}