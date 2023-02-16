func isPalindrome(x int) bool {
    var temp int
    var rev int
    temp = x
    if x < 0 {
        return false
    } 
    for x > 0 {
        dig:= x%10
        rev = rev * 10 + dig
        x = x/10
    }
    if rev == temp {
        return true
    }
    return false
}