func twoSum(nums []int, target int) []int {
	sl:= [] int {}
	for i:=0; i < len(nums); i++ {
		for j:= i+1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				sl = append(sl, i, j)
			}
		}
	}
	return sl
  }