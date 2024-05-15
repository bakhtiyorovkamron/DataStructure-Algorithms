package leetcode

func FindIntersectionValues(nums1 []int, nums2 []int) []int {

	nums1map, nums2map := make(map[int]int, len(nums1)), make(map[int]int, len(nums2))
	a, b := 0, 0
	for i, v := range nums1 {
		nums1map[v] = i
	}
	for i, v := range nums2 {
		nums2map[v] = i
	}
	for _, v := range nums2 {
		if _, ok := nums1map[v]; ok {
			a++
		}
	}
	for _, v := range nums1 {
		if _, ok := nums2map[v]; ok {
			b++
		}
	}
	return []int{b, a}
}
