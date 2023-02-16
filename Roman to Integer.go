func romanToInt(s string) int {
    var res int
   m:= map[rune] int { 'I':1, 'V':5, 'X':10, 'L':50, 'C':100, 'D':500, 'M':1000,}
    for i, val:= range s {
        if i < len(s)-1 && m[val] < m[rune(s[i+1])] {
            res = res - m[val]
        } else {
            res = res + m[val]
        }
    }
    return res
}