var phoneMap = map[rune][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
    if digits == "" {
        return []string{}
    }
    res := []string{""}
    for _, digit := range digits {
        letters := phoneMap[digit]
        temp := make([]string, 0, len(res)*len(letters))
        for _, letter := range letters {
            for _, s := range res {
                temp = append(temp, s+letter)
            }
        }
        res = temp
    }
    return res
}
