package main

func Solve(s string) []int {
  count_A:= 0
  count_a:= 0
  count_dig:= 0
   count_sp:= 0
	sl:= make([]int, 4)
  for _, val:= range s {
    if val>= 97 && val <=122 {
      count_a++
    } else if val >= 65 && val <= 90 {
      count_A++
    } else if val >= '0' && val <= '9' {
      count_dig++
    } else {
      count_sp++
    }
    
  }
  sl[0] = count_A
  sl[1] = count_a
  sl[2] = count_dig
  sl[3] = count_sp
  return sl
}