package main
import "fmt"

func main() {
	var arr1 [6]int  // 数组 自动初始化为全0
	var slice1 []int = arr1[2:5]  // index include 2 3 4 

	fmt.Printf("arr1=%v\n",arr1)
	fmt.Printf("slice1 := arr1[2:5]| %v\n", slice1)
	// 给数组赋值，会改变切片 因为切片是数组的指针
	for i:=0; i<len(arr1); i++ {
		arr1[i] = i
	}
	fmt.Printf("arr1=%v\n", arr1)
	for i:=0; i<len(slice1); i++ {
		fmt.Printf("slice1[%d] = %d\n", i, slice1[i])  // 切片和数组一样可索引
	}
	// see cap and len
	fmt.Printf("The length of arr1 is %d\n", len(arr1)) // 6
	fmt.Printf("The length of slice1 is %d\n", len(slice1))  // 3
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))  // 4
	// 一个切片 s 可以这样扩展到它的大小上限：s = s[:cap(s)]，如果再扩大会导致运行时错误

	fmt.Println("x := []string{\"A\", \"B\", \"c\", \"D\", \"E\", \"F\", \"k\"}")
	x := []string{"A", "B", "c", "D", "E", "F", "k"}
	// 切片 len 7, cap 7
	fmt.Printf("%v len %d\n",x, len(x))
	fmt.Printf("%v cap %d\n",x, cap(x))
	
}

