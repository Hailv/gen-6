package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var result []int

	for i := 0; i < len(nums1); i++ {
		var k int
		isHaveValue := false
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				k = j
				break
			}
		}

		for k < len(nums2) {
			if nums2[k] > nums1[i] {
				result = append(result, nums2[k])
				isHaveValue = true
				break
			}
			k++
		}

		if isHaveValue == false {
			result = append(result, -1)
		}
	}

	return result
}
