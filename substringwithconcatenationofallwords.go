func findSubstring(s string, words []string) []int {
    n := len(s)
    m := len(words)
    wordLen := len(words[0])
    windowLen := m * wordLen
    wordCount := make(map[string]int)
    for _, word := range words {
        wordCount[word]++
    }
    result := []int{}
    for i := 0; i <= n-windowLen; i++ {
        seen := make(map[string]int)
        for j := 0; j < m; j++ {
            idx := i + j*wordLen
            word := s[idx : idx+wordLen]
            if wordCount[word] == 0 {
                break
            }
            seen[word]++
            if seen[word] > wordCount[word] {
                break
            }
            if j == m-1 {
                result = append(result, i)
            }
        }
    }
    return result
}
