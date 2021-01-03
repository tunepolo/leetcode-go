package twosum

// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	complementIndex := make(map[int]int)

	for i, v := range nums {
		complement := target - v
		j, ok := complementIndex[complement]
		if ok {
			return []int{j, i}
		}
		complementIndex[v] = i
	}

	return []int{-1, -1}
}
