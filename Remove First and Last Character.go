package main

func RemoveChar(word string) string {
  str:= ""
  for i:= 1; i <len(word)-1; i++ {
    str = str + string(word[i])
  }
  return str
}