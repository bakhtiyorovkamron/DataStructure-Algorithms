package leetcode

func MySqrt(x int) int {
	var result , mid , end ,begin,v int
	end = x
	for begin <= end {
		mid = (begin+end)/2
		v = mid*mid
		if v == x {
			return mid
		}
		if v > x {
			end = mid-1
		}else{
			result = mid
			begin = mid+1
		}
	}
	return result
}
