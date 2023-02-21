
import "math"
func searchInsert(nums []int, target int) int {
    p1:= 0
    p2:= len(nums)-1
    var mid int
    for p1 <= p2 {
        mid = int(math.Round(float64(p1+p2)/2))
    if nums[mid] == target {
        return mid
    } else if nums[mid] < target {
        p1 = mid + 1
    } else if nums[mid] > target {
        p2 = mid -1
    }
    }
    return p1
}