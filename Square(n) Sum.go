package main

func SquareSum(numbers []int) int {
  var res int
  
  for i:=0; i < len(numbers); i++ {
    res = res + numbers[i] * numbers[i]
  }
  return res
}