package main

import "fmt"

func main() {
	// 声明数组，不指定元素个数，就成了slice
	var arr1 [3]int
	// 自动初始化 对应类型的0值 [0 0 0]
	fmt.Println(arr1)
	fmt.Printf("arr1 type is %T\n", arr1)

	l := len(arr1)
	fmt.Printf("length of arr1 is %d\n", l)

	// 赋值
	for i:=0; i<len(arr1); i++ {
		arr1[i] = i * 2  // 数组元素已经初始化 赋值直接用 = 而不是 :=
	}
	fmt.Println(arr1)
	
	// for range 形式遍历 不要遗漏range关键字
	for index, sth := range arr1 {
		fmt.Printf("index=%d, sth is %v\n", index, sth)
	}  // actually sth is exactly the element arr1[index]

	// 声明数组的时候可以不指定 元素个数，用...来替代，但是要给出元素
	/// var arr2 = [...]string 这样是错误的，既没有指定长度，又没有给出元素
	arr2 := [...]string{"A", "B", "C", "D", "e"}
	fmt.Println(arr2)
	
	// 数组 还可以通过new来创建，只不过返回的是指针(new 返回的都是指针类型)
	var parr3 = new([5]float32)  // parr3 是 *[5]float32
	fmt.Println(parr3)  // &[0, 0, 0, 0, 0]
	fmt.Printf("parr3 type is %T\n", parr3)  // *[5]float32

	// 可以通过指针的 取值运算 得到地址上的值 类似C++的 int* p, *p取值
	var arr3 [5]float32
	arr3 = *parr3
	fmt.Printf("arr3 = *parr3, arr3=%v\n", arr3)

	// 上面的赋值 也可以写作 arr3 := *parr3
	// ## 如果现在修改 arr3, parr3地址上的数组会改变吗？不会
	arr3[2] = 99.25
	fmt.Printf("arr3: %v | parr3 %v\n", arr3, parr3)
	// 在赋值之后，修改arr3 不会导致parr3地址上的数组改变
	/*
	   接上，在函数传参的时候，如果传递的是数组，而不是数组指针，那么只是值传递
	   要在函数内 对原来的数组造成改变，需要传递数组指针
	*/
	fmt.Printf("--------数组作为函数参数----------\n")
	fmt.Printf("在main 函数，arr1 = %v \n",arr1)
	fmt.Printf("在main 函数 arr1地址 %p \n", &arr1)

	testArrPointer(arr1)
	fmt.Printf("运行test之后 main中 arr1=%v 没有改变\n", arr1)

	testArrPointer2(&arr1)
	fmt.Printf("运行test1之后 main中 arr1=%v 发生改变\n", arr1)
}

func testArrPointer(a [3]int) {
	fmt.Printf("在test 函数，a = %v \n",a)
	// 打印参数 a 的地址
	fmt.Printf("在test 函数 a地址 %p \n", &a)
	// 给每一个元素减一
	for i:=0; i<len(a); i++ {
		a[i] -= 1
	}
	// 查看改变后的参数a
	fmt.Printf("test 改变之后a = %v\n", a)
}

func testArrPointer2(a *[3]int) {	
	fmt.Printf("在test1 函数，a = %v \n",a)
	// 打印参数 a 的地址
	fmt.Printf("在test1 函数 a地址 %p \n", a)
	// 给每一个元素减一
	for i:=0; i<len(a); i++ {
		a[i] -= 1
	}
	// 查看改变后的参数a
	fmt.Printf("test1 改变之后a = %v\n", a)
}
