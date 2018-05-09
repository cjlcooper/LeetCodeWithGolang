package main

import (
	"fmt"
	"strconv"  
)

func main() {
	fmt.Println(reverse(901000))
}

func reverse(x int) int {
	str := strconv.Itoa(x)
	reslutStr := ""
	var items []int
	for i := len(str)-1; i >= 0; i-- {
		//delete "0"
		//fmt.Println(str[0:3])
		if (string(str[i])!="0") {
			fmt.Println(i)
			items.append(i)
			//keep "-"
			// if (string(str[0]) == "-") {
			// 	if (i == 0) {
			// 		reslutStr = "-" + reslutStr
			// 	}else {
			// 		reslutStr += string(str[i])
			// 	}
			// }else{
			// 	reslutStr += string(str[i])
			// }
			//fmt.Println(i)
		}
	}
	resultInt,err := strconv.Atoi(reslutStr)  
	if (err != nil) {
		return x
	}

    return resultInt
}