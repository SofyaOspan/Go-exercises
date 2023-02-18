package main

func GetCount(str string) (count int) {
  count = 0
 for _, val:= range str {
   if val == 'a' || val == 'e' || val == 'i' || val == 'o' || val == 'u' {
     count++
   }
 }
  return count
}