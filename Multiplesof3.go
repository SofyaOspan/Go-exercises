

package main

package kata

func Multiple3And5(number int) int {
  // natural * 1 and itself
	//123456789
	//3569 = sum
	//return sum = 23
  if number < 0 {
    return 0
  }
  count:= 0
  for i:= 0; i < number; i++ {
    if i %3 == 0 || i %5 == 0 {
      count = count + int(i)
    }
  }
  return count
}

