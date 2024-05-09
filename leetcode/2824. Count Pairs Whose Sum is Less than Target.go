package leetcode


func CountPairs(nums []int, target int) int {
	counter := 0 
    for i := 0 ; i < len(nums);i++ {
    	for j := i+1; j < len(nums);j++ {
    		if ( nums[i] + nums[j] ) < target {
    			counter++
    		}
    	}
    }
    return counter
}