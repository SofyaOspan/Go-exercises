package kata

func TwoToOne(s1 string, s2 string) string {
    seen := make(map[rune]bool)
    for _, c := range s1 {
        seen[c] = true
    }
    for _, c := range s2 {
        seen[c] = true
    }
    chars := make([]rune, 0, len(seen))
    for c := range seen {
        chars = append(chars, c)
    }
    for i := 0; i < len(chars); i++ {
        for j := i + 1; j < len(chars); j++ {
            if chars[i] > chars[j] {
                chars[i], chars[j] = chars[j], chars[i]
            }
        }
    }
    res := ""
    for _, c := range chars {
        res += string(c)
    }

    return res
}
