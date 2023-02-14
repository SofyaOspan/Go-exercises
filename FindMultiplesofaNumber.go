package main

func FindMultiples(integer, limit int) []int {
 sl:= [] int {}
  for i:= integer; i <= limit; i++ {
    if i % integer == 0 {
      sl = append(sl, i)
    }
  }
  return sl
}