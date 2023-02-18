package main

func BubblesortOnce(numbers []int) []int {
  n:= len(numbers)
  for i:= 0; i < n-1; i++ {
    for j:= 0; j < n-i-1; j++ {
      if numbers[j] > numbers[j+1] {
        numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
}
    }
  }
  return numbers
  }