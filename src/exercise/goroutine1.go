package main

import (
    "fmt"
    "time"	
)

func main() {
	fmt.Println("in main()")
	go longWait()
    go shortWait()
    
    fmt.Println("About to sleep in main()")
    time.Sleep(10 * 1e9) // 10 * 10^9 ns = 10 s	
	fmt.Println("At the end of main()")
}


func longWait() {
	fmt.Println("Beginning longWait()..")
    time.Sleep(5 * 1e9) // 5 * 10^9 ns
	fmt.Println("End of longWait()")
}


func shortWait() {
	fmt.Println("Beginning shortWait()..")
    time.Sleep(2 * 1e9) // 2 * 10^9 ns
	fmt.Println("End of shortWait()")
}

