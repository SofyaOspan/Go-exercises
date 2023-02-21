package main
import "strings"
func ReverseWords(str string) string {  
  words := strings.Split(str, " ")
    for i := 0; i < len(words); i++ {
        runes := []rune(words[i])
        for j, k := 0, len(runes)-1; j < k; j, k = j+1, k-1 {
            runes[j], runes[k] = runes[k], runes[j]
        }
        words[i] = string(runes)
    }
    return strings.Join(words, " ")
}