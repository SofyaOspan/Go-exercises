package kata
import ("fmt"
        "math")
func RowSumOddNumbers(n int) int {
    // your code here
 // 3  = 27 = 3 pow 3
  // we see that n pow 3
  res := math.Pow(float64(n), 3)
  fmt.Println(res)
  return int(res)
}