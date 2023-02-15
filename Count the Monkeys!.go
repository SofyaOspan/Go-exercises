package main

func monkeyCount(n int) []int {
  sl:= make([]int, n)
  for i:= 1; i <= n; i++ {
    sl[i-1] = i
  }
  return sl
}