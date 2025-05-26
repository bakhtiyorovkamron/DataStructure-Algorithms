package leetcode

func ProductExceptSelf(nums []int) []int {
    var result []int
    for i := 0 ; i < len(nums) ; i++ {
    	var product int = 1

    	for j := 0 ; j < len(nums);j++ {
    		if i != j {
    			product *= nums[j]
    		}
    	}
    	result = append(result,product)
    }


    return result
}