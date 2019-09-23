package findmediansortedarrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)
	arrayA, arrayB := nums1, nums2

	// arrayB, mの方が小さいという仮定を置く
	if m > n {
		m, n = len(nums2), len(nums1)
		arrayA, arrayB = nums2, nums1
	}

	iMin, iMax := 0, m
	var i, j int

	for {
		i = (iMin + iMax) / 2
		j = (m+n+1)/2 - i

		if i < iMax && arrayB[j-1] > arrayA[i] {
			iMin = i + 1
		} else if i > iMin && arrayA[i-1] > arrayB[j] {
			iMax = i - 1
		} else {
			// iの探索が完了した
			break
		}
	}

	// 添字のアクセス範囲に注意しながらA・Bの左側の最大値、右側の最小値を確認する
	var median float64
	var leftMax int
	if i == 0 {
		// Aの左が空の場合
		leftMax = arrayB[j-1]
	} else if j == 0 {
		// Bの左がからの場合
		leftMax = arrayA[i-1]
	} else {
		leftMax = max(arrayA[i-1], arrayB[j-1])
	}

	// A・Bの要素数が奇数の時、左側の最大値がmedianである。
	if (m+n)%2 != 0 {
		median = float64(leftMax)
	} else {
		var rightMin int
		if i == m {
			// Aの右が空の場合
			rightMin = arrayB[j]
		} else if j == n {
			// Bの右が空の場合
			rightMin = arrayA[i]
		} else {
			rightMin = min(arrayA[i], arrayB[j])
		}

		median = float64(leftMax+rightMin) / 2.0
	}

	return median
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
