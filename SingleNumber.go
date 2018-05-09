package main

import (
	"fmt"
)

func main(){
	nums := []int{2,2,1}
	//fmt.Println(singelNumer(nums))
	singelNumer(nums)
}

func singelNumer(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		for j := i+1; j < len(nums); j++ {
			fmt.Println(nums[i],nums[j])
		}
	}
	return result
}