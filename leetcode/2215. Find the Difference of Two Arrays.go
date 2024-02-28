package leetcode

func FindDifference(nums1 []int, nums2 []int) [][]int {

	nums1Map, nums2Map := make(map[int]struct{}), make(map[int]struct{})
	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]] = struct{}{}
	}
	for i := 0; i < len(nums2); i++ {
		nums2Map[nums2[i]] = struct{}{}
	}
	result := make([][]int, 2)
	for k, _ := range nums1Map {
		if _, ok := nums2Map[k]; !ok {
			result[0] = append(result[0], k)
		}
	}
	for k, _ := range nums2Map {
		if _, ok := nums1Map[k]; !ok {
			result[1] = append(result[1], k)
		}
	}

	return result
}
