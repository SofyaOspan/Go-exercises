package main

func StringToNumber(str string) int {
  var res int
  var negative bool
  if str[0] == '-' {
    negative = true
    str = str[1:]
  }
  for _,val:= range str {
    res = res*10 + int(val-'0')
    }
  if negative {
    return -res
  }
  return res
}
