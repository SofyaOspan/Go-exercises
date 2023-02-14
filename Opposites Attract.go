package main

func LoveFunc(flower1, flower2 int) bool {
  if flower1 %2 == 0  && flower2 %2 == 1|| flower2 %2 == 0 && flower1 == 1 {
    return true
  }
  return false
}