package main

func BubblesortOnce(numbers []int) []int {
  n:= len(numbers)
  sl:= make([] int, n)
  copy(sl, numbers)
    for j:= 0; j <n-1; j++ {
      if sl[j] > sl[j+1] {
        sl[j], sl[j+1] = sl[j+1], sl[j]
    }
  }
  return sl
  }