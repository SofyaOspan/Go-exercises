package main

func IsPalindrome(str string) bool {
  for i, j := 0, len(str)-1; i < j; i, j = i+1, j-1 {
    if !isAlpha(rune(str[i])) {
      i++
      continue
    }
    if !isAlpha(rune(str[j])) {
      j--
      continue
    }
    if str[i] != str[j] && str[i] != str[j]+32 && str[i]+32 != str[j] {
      return false
    }
  }
  return true
}

func isAlpha(r rune) bool {
  return ('A' <= r && r <= 'Z') || ('a' <= r && r <= 'z')
}