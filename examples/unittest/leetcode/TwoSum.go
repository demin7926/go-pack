package leetcode

func TwoSum(nums []int, target int) []int {
	ret := []int{0, 0}
	//for i in range(0, len(nums)-1) {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ret[0], ret[1] = i, j
				break
			}
		}
	}
	return ret
}
