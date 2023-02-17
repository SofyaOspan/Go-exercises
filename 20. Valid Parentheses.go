func isValid(s string) bool {
    sl:= [] rune {}

    for _, val:= range s {
        if val == '(' || val == '[' || val == '{' {
            sl = append(sl, val)
        } else if val == ')' && len(sl) > 0 && sl[len(sl)-1] == '(' {
            sl= sl[0:len(sl)-1]
        } else if val == ']' && len(sl) > 0 && sl[len(sl)-1] == '[' {
            sl = sl[0:len(sl)-1]
        } else if val == '}' && len(sl) > 0 && sl[len(sl)-1] == '{' {
            sl = sl[0:len(sl)-1]
        } else {
            return false
        }
    }
    return len(sl) == 0
}