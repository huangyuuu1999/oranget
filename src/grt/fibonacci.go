package main

import (
	"fmt"
	"errors"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		ret, _ := Fibonacci(num)
		fmt.Printf("F(%d) = %d\n", num, ret)
	}
}

/*
 1 1 2 3 5 8 11  
*/

func Fibonacci(n int) (int, error) {
	if n <= 0{
		return 0, errors.New("n in Fabonacci(n) must >= 1.")
	}
	if n == 1 || n == 2 {
		return 1, nil
	}else {
		ret1, _ := Fibonacci(n-1)
		ret2, _ := Fibonacci(n-2)
		return ret1 + ret2, nil
	}
}
