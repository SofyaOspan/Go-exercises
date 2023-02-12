package main
import "strings"
func ReverseWords(str string) string {
  words:= strings.Split(str, " ")
  for i,j := 0,len(words)-1; i<j; i,j = i+1, j-1 {
    words[i], words[j] = words[j], words[i]
  }
  return strings.Join(words, " ")
}
//output "The greatest victory is that which requires no battle" --> "battle no requires which that is victory greatest The"