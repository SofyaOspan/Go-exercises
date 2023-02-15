package main

func SumCubes(n int) int {
 var res int
  for i:= 1; i <=n; i++ {
    res = res + i*i*i
  }
  return res
}