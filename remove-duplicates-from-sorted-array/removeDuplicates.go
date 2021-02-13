package removeduplicates

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/

func removeDuplicates(nums []int) int {
	size := len(nums)
	if size < 2 {
		return size
	}

	left, right := 0, 1
	for ; right < size; right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left], nums[right] = nums[right], nums[left]
	}

	return left + 1
}
