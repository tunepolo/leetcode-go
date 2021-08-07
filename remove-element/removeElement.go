package removeElement

// https://leetcode.com/problems/remove-element/

func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v == val {
			continue
		} else {
			nums[index] = v
			index++
		}
	}

	// zero padding
	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}

	return index
}
