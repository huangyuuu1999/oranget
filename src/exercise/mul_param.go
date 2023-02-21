package main

import "fmt"

func main(){
	x :=min(1, 3, 2, 0)
	fmt.Printf("min of %v is %d\n", []int{1, 3, 2, 0} , x)
	slice := []int{7, 9, 3, 5, 4}
	x = min(slice...)
	fmt.Printf("The minimum in the slice is: %d\n", x)
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}