package searchInsertPosition

// https://leetcode.com/problems/search-insert-position/

// O(log n)
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		pos := (left + right) / 2

		if nums[pos] == target {
			return pos
		} else if nums[pos] > target {
			right = pos - 1
		} else {
			left = pos + 1
		}
	}

	return left
}

// O(n)
func searchInsert_Bruteforce(nums []int, target int) int {
	for i, v := range nums {
		if v >= target {
			return i
		}
	}

	return len(nums)
}
