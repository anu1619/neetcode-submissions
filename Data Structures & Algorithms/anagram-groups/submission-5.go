func groupAnagrams(strs []string) [][]string {
    groups := make(map[[26]int][]string)

    for _, str := range strs {
        var freq [26]int

        for i := 0; i < len(str); i++ {
            freq[str[i]-'a']++
        }

        groups[freq] = append(groups[freq], str)
    }

    result := make([][]string, 0, len(groups))

    for _, group := range groups {
        result = append(result, group)
    }

    return result
}