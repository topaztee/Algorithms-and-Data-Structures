package main

import (
	"fmt"
)

func main() {
	nums := []int{9,8,6,3,6,1,2}
	fmt.Println(mergeSort(nums))
}

func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	half := len(nums) / 2
	l := mergeSort(nums[:half])
	r := mergeSort(nums[half:])
	return merge(l,r)
}

func merge(a, b []int) []int {
	ret := make([]int, 0, len(a)+len(b))

	for len(a) > 0 || len(b) > 0 {
		if(len(a) == 0) {
			return append(ret, b...)
		}
		if(len(b) == 0) {
			return append(ret, a...)
		}
		if(a[0] <= b[0]) {
			ret = append(ret,a[0])
			a = a[1:]
		} else {
			ret = append(ret, b[0])
			b = b[1:]
		}
	}
	return ret
}

half := len(nums)/2
l := merge(num[:half])

return merge(l,r)