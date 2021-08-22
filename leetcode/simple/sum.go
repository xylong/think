package simple

// TwoSum 两数之和
// @param nums 整数数组
// @param target 整数目标值
// @link https://leetcode-cn.com/problems/two-sum
func TwoSum(nums []int, target int) []int {
	length := len(nums)
	index := make([]int, 0)

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if nums[i]+nums[j] == target {
				index = append(index, i, j)
			}
		}
	}

	return index
}
