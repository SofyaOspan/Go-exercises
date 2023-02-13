package main

func PowersOfTwo(n int) []uint64 {
  res:= [] uint64{1}
  for i:= 1; i <=n; i++ {
    res = append(res, 2* res[i-1])
  }
  return res
}