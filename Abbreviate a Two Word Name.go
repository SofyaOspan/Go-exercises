package main

func AbbrevName(name string) string {
  str:= ""
  for i, val:= range name {
    if val == ' ' {
      str = str + string(name[0]) + "." + string(name[i+1])
    }
  }
  return str
}