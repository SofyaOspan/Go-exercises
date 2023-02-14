package main

func ToAlternatingCase(str string) string {
  res := ""
  for _, val:= range str  { 
  switch {
    case val >= 65 && val <= 90:
      res = res + string(val + 32)
    case val >= 97 && val <= 122:
    res = res + string(val-32)
    case val == ' ' :
    res = res + string(' ')
    case val >= '0' && val<= '9':
    res = res + string(val)
    case val == '.' :
    res = res + string('.')
    case val == '<' :
    res = res + string('<')
    case val == '=':
    res = res + string('=')
    case val == '>':
    res = res + string('>')
  }
  }
  return res
}