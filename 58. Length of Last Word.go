func lengthOfLastWord(s string) int {
    count:= 0
    flag := false
    
    for i:= len(s)-1; i >= 0; i -- {
        if s[i] == ' ' {
            if flag {
                break
            } else {
            continue
            }
        }
        flag = true 
             count++
    }
    return count
}