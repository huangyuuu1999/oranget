package main

import "fmt"

func main() {
	arr1 := [...]float32{1.1, 2.3, 3.5, 4.8, 2.34, 3.14}  // 创建数组
	slice1 := arr1[1:3]  // slice1 type []float32 是切片
	fmt.Printf(" arr1 %v\n slice1 := arr1[1:3] | %v\n", arr1, slice1)

	fmt.Printf(" len(slice1) %v\n len(arr1) %v\n cap(slice1) %v\n", len(slice1), len(arr1), cap(slice1))
	fmt.Println("+-------------------------+")
	fmt.Println("| now, scale slice1 here. |")
	fmt.Println("+-------------------------+")
	slice1 = slice1[:cap(slice1)]
	fmt.Println("+-------------------------+")
	fmt.Printf(" slice1 len after scale %v\n", len(slice1))
	fmt.Println("+-------------------------+")
	fmt.Printf(" slice1 cap after scale %v\n", cap(slice1))
	fmt.Println("+-------------------------+")
}
