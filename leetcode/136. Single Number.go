package leetcode

import "fmt"

//4, 1, 2, 1, 2

func SingleNumber(nums []int) int {
	mymap := make(map[int]int)
	totnumber := 0
	for _, v := range nums {
		if _, ok := mymap[v]; ok {
			mymap[v] += 1
		} else {
			mymap[v] = 1
		}
	}
	for i, _ := range mymap {
		if mymap[i] == 1 {

			totnumber = i
		}
		fmt.Println(i, mymap[i])
	}
	return totnumber
}
