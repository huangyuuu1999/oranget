package main

import "fmt"


func main() {
	// 切片是数组 或者数组局部的引用，相当于指针
	var arr1 [6]int  //这是一个数组 是值类型, 自动初始化0
	var slice1 []int = arr1[2:5]  // 指向数组的第 2,3,4 下标元素的切片
	fmt.Printf("slice1 type: %T\n", slice1)
	for a, b := range slice1 {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}
	fmt.Printf("slice1 len %d\n", len(slice1))
	// 也可以取得数组的全部元素
	slice2 := arr1[:]  // 相当于[0:len(arr1)]
	fmt.Printf("slice2 = arr1[:] => %v\n", slice2)

	// 改变数组 arr1 slice 的内容会改变吗？
	arr1[1] = 9
	arr1[2] = 8
	arr1[3] = 7
	arr1[4] = 6
	arr1[5] = 4
	fmt.Printf("change arr1, would slice1 change?\n")
	fmt.Printf("changed arr1 %v\n", arr1)
	fmt.Printf("changed slice1 %v\n", slice1)
	// 切片还可以继续切
	slice2 = slice1[1:3]
	fmt.Printf("slice2=slice1[1:3], cut from slice1: %v\n",slice2)

	// 切片还可以更简洁地产生
	slc := [3]int{1, 2, 3}
	slc2 := []int{1, 2, 3, 4}
	fmt.Printf("slc := [3]int{1, 2, 3}| %v\n", slc)
	fmt.Printf("slc2 := []int{1, 2, 3, 4}| %v\n", slc2)
	// 3. cap 函数 len 函数的区别
	arr3 := [7]int{3, 4, 5, 6, 7, 8, 9}
	// slice3 := arr3[:-1] 不能是负数 与 Python 不同
	slice3 := arr3[:len(arr3)-1] // 去掉最后一个元素
	fmt.Printf("arr3=%v  slice3=%v\n", arr3, slice3)
}
