package findMedianSortedArrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	array_a, array_b := nums1, nums2

	// array_b, mの方が小さいという仮定を置く
	if m > n {
		m, n = len(nums2), len(nums1)
		array_a, array_b = nums2, nums1
	}

	i_min, i_max := 0, m
	var i, j int

	for {
		i = (i_min + i_max) / 2
		j = (m+n+1)/2 - i

		if i < i_max && array_b[j-1] > array_a[i] {
			i_min = i + 1
		} else if i > i_min && array_a[i-1] > array_b[j] {
			i_max = i - 1
		} else {
			// iの探索が完了した
			break
		}
	}

	// 添字のアクセス範囲に注意しながらA・Bの左側の最大値、右側の最小値を確認する
	var median float64
	var left_max int
	if i == 0 {
		// Aの左が空の場合
		left_max = array_b[j-1]
	} else if j == 0 {
		// Bの左がからの場合
		left_max = array_a[i-1]
	} else {
		left_max = max(array_a[i-1], array_b[j-1])
	}

	// A・Bの要素数が奇数の時、左側の最大値がmedianである。
	if (m+n)%2 != 0 {
		median = float64(left_max)
	} else {
		var right_min int
		if i == m {
			// Aの右が空の場合
			right_min = array_b[j]
		} else if j == n {
			// Bの右が空の場合
			right_min = array_a[i]
		} else {
			right_min = min(array_a[i], array_b[j])
		}

		median = float64(left_max+right_min) / 2.0
	}

	return median
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
