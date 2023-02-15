package kata
func SumOfIntegersInString(strng string) int {
    var res int
    var curr int
    for _, val := range strng {
        if isAlpha(val) {
            curr = curr*10 + int(val-'0')
        } else {
            res += curr
            curr = 0
        }
    }
    res += curr 
    return res
}

func isAlpha(r rune) bool {
    return r >= '0' && r <= '9'
}