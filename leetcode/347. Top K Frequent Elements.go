package leetcode

func TopKFrequent(nums []int, k int) []int {
	f := make(map[int]int)
	for _, v := range nums {
		f[v]++
	}
	bucket := make([][]int, len(nums)+1)
	for k, v := range f {
		bucket[v] = append(bucket[v], k)
	}
	res := make([]int, 0)
	for i := len(bucket) - 1; i >= 0; i-- {
		if len(bucket[i]) > 0 {
			res = append(res, bucket[i]...)
		}
		if len(res) >= k {
			break
		}
	}
	return res[:k]
}
