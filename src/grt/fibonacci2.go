package main

import "fmt"

func main(){
	x := [3]int{2, 6}
	fmt.Println(x) 
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range(a){
		fmt.Printf("F[%d] = %d\n", v, Fibo(v))
	}

	TestNotFullArr()
}

func Fibo(n int) int {
	// <0 的判断 最好 是需要的
	if n == 1 || n == 2 {
		return 1
	}else {
		v1 := Fibo(n-1)
		v2 := Fibo(n-2)
		return v1 + v2
	}
}

func TestNotFullArr() {
	arr := [5]int{1, 2, 3} // [5]int 类型数组，只给赋三个值 会出错吗？
    for index,v := range arr{
	    fmt.Printf("arr[%d]=%d\n" ,index ,v) //会自动将缺失的值赋为0
	}
	/*
	 自动赋值为0，在上面的 main 函数调用 Fibo 的时候 会出问题
	 若只是给[10]int 的a 初始化9个值，会在最后一位补充0，而0作为参数传入Fibo的时候会报错	
	*/
}
