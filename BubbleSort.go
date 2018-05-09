package main 

import (
	"fmt"
)

func main() {
	values := []int{2,44,21,4,6,33,1,9,223}
	fmt.Println(values)
	BubbleSort(values)
}

//sort
func BubbleSort(values []int) {
	for i := 0; i < len(values)-1; i++ {
		for j := i+1; j < len(values); j++ {
			if values[i]>values[j] {
				values[i],values[j] = values[j],values[i]
			}
		}
	}
	fmt.Println(values)
}