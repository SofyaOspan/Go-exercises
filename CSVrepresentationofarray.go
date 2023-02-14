package main

import "fmt"
func ToCsvText(array [][]int) string {
  res := ""
  for i, val := range array {
    for j, jval := range val {
      res += fmt.Sprintf("%d", jval)
      if j != len(val) - 1 {
        res += ","
      }
    }
    if i != len(array) - 1 {
      res += "\n"
    }
  }
  return res
}