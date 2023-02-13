package main

func Grow(arr []int) int{
 res:= 1
  for _, val:= range arr {
    res = res * val
  }
  return res 
}