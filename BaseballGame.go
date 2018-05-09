package main

import (
	"fmt"
	"strconv"
)

func  main() {
	ops := []string{"5","2","C","D","+"}
	fmt.Println(calPoints(ops))
}

func calPoints(ops []string) int {
	resultPoint := 0
	for i := 0; i < len(ops); i++ {
		fmt.Println(string(ops[i]))
		if (string(ops[i]) != "C" && string(ops[i]) != "D") {
			//add score
			singlePoint,err := strconv.Atoi(string(ops[i]))
			if (nil == err) {
				resultPoint += singlePoint
			}
		}else{
			if (string(ops[i]) == "C") {
				//delete score
				singlePoint,err := strconv.Atoi(string(ops[i-1]))
				if (nil == err) {
					resultPoint -= singlePoint
				}
			}else if (string(ops[i]) == "D"){
				//x2
				resultPoint = resultPoint*2
			}
		}
	}
	return resultPoint
}